package swagger

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"

	"github.com/getkin/kin-openapi/openapi3"
	converter2 "github.com/idsulik/swama/internal/converter"
	"github.com/idsulik/swama/internal/printer"
	"github.com/olekukonko/tablewriter"
)

const (
	GroupByTag    = "tag"
	GroupByMethod = "method"
	GroupByNone   = "none"
)

type Endpoints interface {
	ListEndpoints(method, endpoint, tag, group string) error
	ViewEndpoint(method, endpoint string) error
	ConvertEndpoint(method, endpoint, toType string) error
}

type endpoints struct {
	doc *openapi3.T
}

func NewEndpoints(doc *openapi3.T) Endpoints {
	return &endpoints{
		doc: doc,
	}
}

// ListEndpoints lists all available API endpoints in the Swagger/OpenAPI file.
func (e *endpoints) ListEndpoints(method, endpoint, tag, group string) error {
	type groupItem struct {
		method  string
		path    string
		summary string
		tags    string
	}
	groupedEndpoints := make(map[string][]groupItem)
	for _, path := range e.doc.Paths.InMatchingOrder() {
		for m, operation := range e.doc.Paths.Find(path).Operations() {
			if endpoint != "" {
				if matched, _ := regexp.MatchString(fmt.Sprintf("^%s$", endpoint), path); !matched {
					continue
				}
			}

			if method != "" && m != method {
				continue
			}

			if tag != "" && !slices.Contains(operation.Tags, tag) {
				continue
			}

			tags := ""
			if len(operation.Tags) > 0 {
				tags = fmt.Sprintf("%v", operation.Tags)
			}

			if group != "" {
				keys := make([]string, 0)
				if group == GroupByTag {
					for _, tagName := range operation.Tags {
						tag := e.doc.Tags.Get(tagName)
						key := tagName
						if tag != nil {
							key += fmt.Sprintf(" (%s)", tag.Description)
						}
						keys = append(keys, key)
					}
				} else if group == GroupByMethod {
					keys = append(keys, m)
				} else {
					keys = append(keys, "none")
				}

				for _, key := range keys {
					if _, ok := groupedEndpoints[key]; !ok {
						groupedEndpoints[key] = make([]groupItem, 0)
					}
					groupedEndpoints[key] = append(
						groupedEndpoints[key],
						groupItem{
							method:  m,
							path:    path,
							summary: operation.Summary,
							tags:    tags,
						},
					)
				}
				continue
			}
		}
	}

	// Sort and print the grouped endpoints
	sortedKeys := slices.Sorted(maps.Keys(groupedEndpoints))
	var table *tablewriter.Table
	fmt.Println()
	for _, key := range sortedKeys {
		if key != GroupByNone {
			fmt.Printf("%s\n", key)
		}

		table = tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetRowLine(true)
		table.SetHeader([]string{"Method", "Path", "Summary", "Tags"})
		values := groupedEndpoints[key]
		for _, value := range values {
			methodColor := tablewriter.FgWhiteColor
			if value.method == "GET" {
				methodColor = tablewriter.FgBlueColor
			} else if value.method == "POST" {
				methodColor = tablewriter.FgGreenColor
			} else if value.method == "PUT" {
				methodColor = tablewriter.FgYellowColor
			} else if value.method == "DELETE" {
				methodColor = tablewriter.FgRedColor
			}

			table.Rich(
				[]string{value.method, value.path, value.summary, value.tags}, []tablewriter.Colors{
					{methodColor},
				},
			)
		}

		table.Render()

		if key != GroupByNone {
			fmt.Println()
		}
	}

	return nil
}

// ViewEndpoint shows details about a specific API endpoint.
func (e *endpoints) ViewEndpoint(method, endpoint string) error {
	operation, err := e.findOperation(method, endpoint)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Method", "Path", "Summary", "Tags"})
	table.Append([]string{method, endpoint, operation.Summary, fmt.Sprintf("%v", operation.Tags)})

	table.Render()

	if len(operation.Parameters) > 0 {
		fmt.Println("Parameters:")
		printer.PrintParameters(operation)
	}

	if operation.RequestBody != nil {
		fmt.Println("Request Body:")
		printer.PrintRequestBody(operation)
	}

	if operation.Responses != nil {
		fmt.Println("Responses:")
		printer.PrintResponses(operation)
	}

	return nil
}

// ConvertEndpoint converts an endpoint to curl or fetch.
func (e *endpoints) ConvertEndpoint(method, endpoint, toType string) error {
	operation, err := e.findOperation(method, endpoint)

	if err != nil {
		return err
	}

	converter, err := converter2.NewConverter(toType)

	if err != nil {
		return err
	}

	converted := converter.ConvertEndpoint(method, endpoint, operation)
	fmt.Println(converted)

	return nil
}

func (e *endpoints) findOperation(method, endpoint string) (*openapi3.Operation, error) {
	path := e.doc.Paths.Find(endpoint)

	if path != nil {
		for m, operation := range path.Operations() {
			if method != "" && m != method {
				return operation, nil
			}
		}
	}

	return nil, fmt.Errorf("endpoint %s not found", endpoint)
}
