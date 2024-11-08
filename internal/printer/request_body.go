package printer

import (
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/olekukonko/tablewriter"
)

func PrintRequestBody(requestBody *openapi3.RequestBody) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Type", "Properties"})
	propertiesToContentTypes := make(map[string][]string)
	for contentType, content := range requestBody.Content {
		properties := strings.Join(getProperties(content.Schema.Value), "\n")
		propertiesToContentTypes[properties] = append(propertiesToContentTypes[properties], contentType)
	}

	for properties, contentTypes := range propertiesToContentTypes {
		table.Append([]string{strings.Join(contentTypes, "\n"), properties})
	}
	table.Render()
}
