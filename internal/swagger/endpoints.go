package swagger

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	converter2 "github.com/idsulik/swama/internal/converter"
	"github.com/idsulik/swama/internal/model"
	"github.com/idsulik/swama/internal/printer"
	"github.com/idsulik/swama/internal/util"
	"github.com/olekukonko/tablewriter"
)

const (
	GroupByTag    = "tag"
	GroupByMethod = "method"
	GroupByNone   = "none"
)

type ListOptions struct {
	Method   string
	Endpoint string
	Tag      string
	Group    string
}

type ViewOptions struct {
	Method   string
	Endpoint string
}

type ConvertOptions struct {
	Method   string
	Endpoint string
	ToType   string
}

type Endpoints interface {
	ListEndpoints(options ListOptions) error
	ViewEndpoint(options ViewOptions) error
	ConvertEndpoint(options ConvertOptions) error
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
func (e *endpoints) ListEndpoints(options ListOptions) error {
	type groupItem struct {
		method  string
		path    string
		summary string
		tags    string
	}
	groupedEndpoints := make(map[string][]groupItem)
	for _, path := range e.doc.Paths.InMatchingOrder() {
		for m, operation := range e.doc.Paths.Find(path).Operations() {
			if options.Endpoint != "" {
				if matched, _ := regexp.MatchString(fmt.Sprintf("^%s$", options.Endpoint), path); !matched {
					continue
				}
			}

			if options.Method != "" && m != options.Method {
				continue
			}

			if options.Tag != "" && !slices.Contains(operation.Tags, options.Tag) {
				continue
			}

			tags := ""
			if len(operation.Tags) > 0 {
				tags = fmt.Sprintf("%v", operation.Tags)
			}

			if options.Group != "" {
				keys := make([]string, 0)
				if options.Group == GroupByTag {
					for _, tagName := range operation.Tags {
						tag := e.doc.Tags.Get(tagName)
						key := tagName
						if tag != nil {
							key += fmt.Sprintf(" (%s)", tag.Description)
						}
						keys = append(keys, key)
					}
				} else if options.Group == GroupByMethod {
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
		sort.Slice(
			values, func(i, j int) bool {
				if values[i].method != values[j].method {
					return values[i].method > values[j].method
				}
				return values[i].path < values[j].path
			},
		)
		for _, value := range values {
			table.Rich(
				[]string{value.method, value.path, value.summary, value.tags}, []tablewriter.Colors{
					{util.GetMethodColor(value.method)},
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
func (e *endpoints) ViewEndpoint(options ViewOptions) error {
	operation, err := e.findOperation(options.Method, options.Endpoint)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Method", "Path", "Summary", "Tags"})
	table.Rich(
		[]string{operation.Method, operation.Path, operation.Summary, fmt.Sprintf("%v", operation.Tags)},
		[]tablewriter.Colors{{util.GetMethodColor(operation.Method)}},
	)

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
func (e *endpoints) ConvertEndpoint(options ConvertOptions) error {
	operation, err := e.findOperation(options.Method, options.Endpoint)

	if err != nil {
		return err
	}

	converter, err := converter2.NewConverter(options.ToType)

	if err != nil {
		return err
	}

	converted := converter.ConvertEndpoint(options.Method, options.Endpoint, operation)
	fmt.Println(converted)

	return nil
}

func (e *endpoints) findOperation(method, endpoint string) (*model.Operation, error) {
	path := e.doc.Paths.Find(endpoint)

	if path != nil {
		for m, operation := range path.Operations() {
			if method != "" && strings.ToLower(m) == strings.ToLower(method) {
				return model.NewOperation(m, endpoint, operation), nil
			}
		}
	}

	return nil, fmt.Errorf("endpoint %s not found", endpoint)
}
