package printer

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/olekukonko/tablewriter"
)

func PrintParameters(operation *openapi3.Operation) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetHeader([]string{"In", "Parameter", "Type", "Required", "Description"})
	for _, p := range operation.Parameters {
		value := p.Value
		description := "-"
		if value.Description != "" {
			description = value.Description
		}
		table.Append(
			[]string{
				value.In,
				value.Name,
				enrichPropertyName(value.Name, value.Schema),
				fmt.Sprintf("%v", value.Required),
				description,
			},
		)
	}
	table.Render()
}
