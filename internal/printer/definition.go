package printer

import (
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func PrintDefinition(name string, definition map[string]interface{}) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetHeader([]string{"Name", "Type", "Properties"})
	t := definition["type"].(string)
	properties := slices.Sorted(maps.Keys(definition["properties"].(map[string]interface{})))

	table.Append([]string{name, t, strings.Join(properties, ", ")})

	table.Render()
}
