package swagger

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/idsulik/swama/internal/printer"
	"github.com/olekukonko/tablewriter"
)

type Components interface {
	ListComponents() error
	ViewComponent(name string) error
}

type components struct {
	doc *openapi3.T
}

func NewComponents(doc *openapi3.T) Components {
	return &components{
		doc: doc,
	}
}

// ListComponents lists all available API components in the Swagger/OpenAPI file.
func (e *components) ListComponents() error {
	var sortedNames []string
	if e.doc.Components == nil {
		sortedNames = slices.Sorted(maps.Keys(e.doc.Extensions["definitions"].(map[string]interface{})))
	} else {
		sortedNames = slices.Sorted(maps.Keys(e.doc.Components.Schemas))
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)
	table.SetHeader([]string{"Name", "Type", "Description"})
	for _, name := range sortedNames {
		types := "-"
		description := "-"
		if e.doc.Components == nil {
			definitions := e.doc.Extensions["definitions"].(map[string]interface{})
			definition := definitions[name].(map[string]interface{})
			if definition["type"] != nil {
				types = definition["type"].(string)
			}
			if definition["description"] != nil {
				description = definition["description"].(string)
			}
		} else {
			schema := e.doc.Components.Schemas[name]
			if schema.Value.Description != "" {
				description = schema.Value.Description
			}
			if schema.Value.Type != nil {
				types = strings.Join(schema.Value.Type.Slice(), ", ")
			}
		}

		table.Append([]string{name, types, description})
	}
	table.Render()
	return nil
}

// ViewComponent shows details about a specific API component.
func (e *components) ViewComponent(name string) error {
	if e.doc.Components != nil {
		for n, schema := range e.doc.Components.Schemas {
			if strings.ToLower(n) == strings.ToLower(name) {
				printer.PrintSchema(n, schema.Value)
				return nil
			}
		}
	}
	definitions, found := e.doc.Extensions["definitions"]
	if found {
		for n, definition := range definitions.(map[string]interface{}) {
			if strings.ToLower(n) == strings.ToLower(name) {
				printer.PrintDefinition(n, definition.(map[string]interface{}))
				return nil
			}
		}
	}
	fmt.Printf("Component %s not found\n", name)
	return nil
}
