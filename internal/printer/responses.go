package printer

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/olekukonko/tablewriter"
)

func PrintResponses(operation *openapi3.Operation) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetHeader([]string{"Name", "Content", "Description"})
	sortedCodes := slices.Sorted(maps.Keys(operation.Responses.Map()))
	for _, code := range sortedCodes {
		response := operation.Responses.Value(code)
		content := "-"
		if response.Value.Content != nil {
			propertiesToContentTypes := make(map[string][]string)
			for contentType := range response.Value.Content {
				properties := strings.Join(getProperties(response.Value.Content[contentType].Schema.Value), "\n")
				propertiesToContentTypes[properties] = append(propertiesToContentTypes[properties], contentType)
			}

			for properties, contentTypes := range propertiesToContentTypes {
				content = fmt.Sprintf("types:\n%s\n\nproperties:\n%s", strings.Join(contentTypes, "\n"), properties)
			}
		}

		description := "-"
		if response.Value.Description != nil {
			description = *response.Value.Description
		}

		table.Append([]string{code, content, description})
	}

	table.Render()
}
