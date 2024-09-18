package printer

import (
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
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.SetHeader([]string{"Name", "Content Types", "Properties", "Description"})
	sortedCodes := slices.Sorted(maps.Keys(operation.Responses.Map()))
	for _, code := range sortedCodes {
		response := operation.Responses.Value(code)
		description := "-"
		if response.Value.Description != nil {
			description = *response.Value.Description
		}

		if response.Value.Content != nil {
			propertiesToContentTypes := make(map[string][]string)
			for contentType := range response.Value.Content {
				properties := strings.Join(getProperties(response.Value.Content[contentType].Schema.Value), "\n")
				propertiesToContentTypes[properties] = append(propertiesToContentTypes[properties], contentType)
			}

			for properties, contentTypes := range propertiesToContentTypes {
				table.Append([]string{code, strings.Join(contentTypes, "\n"), properties, description})
			}
		} else {
			table.Append([]string{code, "-", "-", description})
		}

	}

	table.Render()
}
