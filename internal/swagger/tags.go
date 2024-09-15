package swagger

import (
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/olekukonko/tablewriter"
)

type Tags interface {
	ListTags() error
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
func (e *tags) ListTags() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Name", "Description"})
	for _, tag := range e.doc.Tags {
		table.Append([]string{tag.Name, tag.Description})
	}

	table.Render()

	return nil
}
