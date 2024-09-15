package swagger

import (
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/olekukonko/tablewriter"
)

type Servers interface {
	ListServers() error
}

type servers struct {
	doc *openapi3.T
}

func NewServers(doc *openapi3.T) Servers {
	return &servers{
		doc: doc,
	}
}

// ListServers lists all available servers in the Swagger/OpenAPI file.
func (e *servers) ListServers() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"URL", "Description"})
	for _, server := range e.doc.Servers {
		table.Append([]string{server.URL, server.Description})
	}

	table.Render()

	return nil
}
