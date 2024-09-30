package printer

import (
	"os"
	"strings"

	"github.com/idsulik/swama/internal/model"
	"github.com/olekukonko/tablewriter"
)

func PrintRequestBody(operation *model.Operation) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Type", "Properties"})
	propertiesToContentTypes := make(map[string][]string)
	for contentType, content := range operation.RequestBody.Value.Content {
		properties := strings.Join(getProperties(content.Schema.Value), "\n")
		propertiesToContentTypes[properties] = append(propertiesToContentTypes[properties], contentType)
	}

	for properties, contentTypes := range propertiesToContentTypes {
		table.Append([]string{strings.Join(contentTypes, "\n"), properties})
	}
	table.Render()
}
