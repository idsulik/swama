package cmd

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"

	"github.com/idsulik/swama/pkg/openapi"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type ListConfig struct {
	method   string
	endpoint string
	tag      string
	group    string
}

const (
	GroupByTag    = "tag"
	GroupByMethod = "method"
	GroupByNone   = "none"
)

// Command-specific flags for the list command
var listConfig = ListConfig{}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all API endpoints from a Swagger file",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		swagger, err := openapi.ParseSwaggerFile(ctx, config.SwaggerPath)
		if err != nil {
			return fmt.Errorf("error parsing Swagger file: %w", err)
		}

		type groupItem struct {
			method  string
			path    string
			summary string
			tags    string
		}
		groupedEndpoints := make(map[string][]groupItem)
		for _, path := range swagger.Paths.InMatchingOrder() {
			for method, operation := range swagger.Paths.Find(path).Operations() {
				if listConfig.endpoint != "" {
					if matched, _ := regexp.MatchString(fmt.Sprintf("^%s$", listConfig.endpoint), path); !matched {
						continue
					}
				}

				if listConfig.method != "" && method != listConfig.method {
					continue
				}

				if listConfig.tag != "" && !slices.Contains(operation.Tags, listConfig.tag) {
					continue
				}

				tags := ""
				if len(operation.Tags) > 0 {
					tags = fmt.Sprintf("%v", operation.Tags)
				}

				if listConfig.group != "" {
					keys := make([]string, 0)
					if listConfig.group == GroupByTag {
						for _, tag := range operation.Tags {
							description := swagger.Tags.Get(tag).Description
							keys = append(keys, fmt.Sprintf("%s (%s)", tag, description))
						}
					} else if listConfig.group == GroupByMethod {
						keys = append(keys, method)
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
								method:  method,
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
			table.SetHeader([]string{"Method", "Path", "Summary", "Tags"})
			values := groupedEndpoints[key]
			for _, value := range values {
				table.Append([]string{value.method, value.path, value.summary, value.tags})
			}

			table.Render()

			if key != GroupByNone {
				fmt.Println()
			}
		}

		return nil
	},
}

func init() {
	listCmd.Flags().StringVarP(&listConfig.endpoint, "endpoint", "e", "", "Filter by endpoint, supports wildcard")
	listCmd.Flags().StringVarP(&listConfig.method, "method", "m", "", "Filter by method")
	listCmd.Flags().StringVarP(&listConfig.tag, "tag", "t", "", "Filter by tag")
	listCmd.Flags().StringVarP(&listConfig.group, "group", "g", GroupByTag, "Group output by tag, method")
}
