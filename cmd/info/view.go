package info

import (
	"fmt"
	"os"

	"github.com/idsulik/swama/cmd/config"
	"github.com/idsulik/swama/internal/swagger"
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
	table.Append([]string{"Description", doc.Info.Description})
	table.Append([]string{"Terms of Service", doc.Info.TermsOfService})
	table.Append([]string{"Contact", doc.Info.Contact.Name})
	table.Append([]string{"License", doc.Info.License.Name})

	table.Render()

	return nil
}
