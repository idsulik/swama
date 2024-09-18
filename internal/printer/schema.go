package printer

import (
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/olekukonko/tablewriter"
)

func PrintSchema(name string, schema *openapi3.Schema) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetHeader([]string{"Name", "Type", "Properties", "Description"})

	types := "-"
	description := "-"
	if schema.Description != "" {
		description = schema.Description
	}
	if schema.Type != nil {
		types = strings.Join(schema.Type.Slice(), ", ")
	}

	table.Append(
		[]string{
			name,
			types,
			strings.Join(getProperties(schema), "\n"),
			description,
		},
	)

	table.Render()
}
