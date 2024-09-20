package swagger

import (
	"fmt"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/olekukonko/tablewriter"
)

type Tags interface {
	ListTags() error
	ViewTag(name string) error
}

type tags struct {
	doc *openapi3.T
}

func NewTags(doc *openapi3.T) Tags {
	return &tags{
		doc: doc,
	}
}

// ListTags lists all available tags in the Swagger/OpenAPI file.
func (t *tags) ListTags() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Name", "Description", "External Docs"})

	for _, tag := range t.doc.Tags {
		t.printTagDetails(table, tag)
	}

	table.Render()

	return nil
}

// ViewTag shows details about a specific API tag.
func (t *tags) ViewTag(name string) error {
	for _, tag := range t.doc.Tags {
		if strings.ToLower(name) == strings.ToLower(tag.Name) {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetAutoWrapText(false)
			table.SetHeader([]string{"Name", "Description", "External Docs"})

			t.printTagDetails(table, tag)
			table.Render()
			return nil
		}
	}

	return fmt.Errorf("tag not found")
}

func (t *tags) printTagDetails(table *tablewriter.Table, tag *openapi3.Tag) {
	externalDocs := "-"
	if tag.ExternalDocs != nil {
		externalDocs = fmt.Sprintf("%s (%s)", tag.ExternalDocs.Description, tag.ExternalDocs.URL)
	}
	table.Append([]string{tag.Name, tag.Description, externalDocs})
}
