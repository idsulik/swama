package printer

import (
	"fmt"
	"os"

	"github.com/idsulik/swama/internal/model"
	"github.com/olekukonko/tablewriter"
)

func PrintParameters(operation *model.Operation) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetHeader([]string{"In", "Parameter", "Type", "Required", "Description"})
	for _, p := range operation.Parameters {
		description := "-"
		if p.Description != "" {
			description = p.Description
		}

		table.Append(
			[]string{
				p.In,
				p.Name,
				p.Type,
				fmt.Sprintf("%v", p.Required),
				description,
			},
		)
	}
	table.Render()
}
