package info

import (
	"fmt"
	"os"

	"github.com/idsulik/swama/v2/cmd/config"
	"github.com/idsulik/swama/v2/internal/swagger"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// vewViewCommand creates the "view" subcommand
func vewViewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "Displays information about the Swagger file",
		RunE:  viewCommandFunc,
	}

	return cmd
}

func viewCommandFunc(cmd *cobra.Command, _ []string) error {
	doc, err := swagger.LoadSwaggerFile(cmd.Context(), config.SwaggerPath)

	if err != nil {
		return fmt.Errorf("failed to load Swagger file: %w", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	table.Append([]string{"Title", doc.Info.Title})
	table.Append([]string{"Version", doc.Info.Version})
	if doc.Info.Contact != nil {
		if doc.Info.Contact.Email != "" {
			table.Append([]string{"Email", doc.Info.Contact.Email})
		}
	}
	if doc.Info.License != nil {
		if doc.Info.License.Name != "" {
			table.Append([]string{"License", fmt.Sprintf("%s (%s)", doc.Info.License.Name, doc.Info.License.URL)})
		}
	}
	table.Append([]string{"Description", doc.Info.Description})
	table.Append([]string{"Terms of Service", doc.Info.TermsOfService})

	if doc.ExternalDocs != nil {
		table.Append([]string{"External docs", doc.ExternalDocs.URL})
	}
	table.Render()

	return nil
}
