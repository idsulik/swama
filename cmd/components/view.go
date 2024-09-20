package components

import (
	"fmt"

	"github.com/idsulik/swama/cmd/config"
	"github.com/idsulik/swama/internal/swagger"
	"github.com/spf13/cobra"
)

type viewConfig struct {
	name string
}

// Command-specific flags for the view command
var viewCfg = viewConfig{}

// newViewCommand creates the "components view" subcommand
func newViewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "View details of a specific component",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(cmd.Context(), config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			components := swagger.NewComponents(doc)

			return components.ViewComponent(viewCfg.name)
		},
	}

	cmd.Flags().StringVarP(&viewCfg.name, "name", "n", "", "Name of the component to view")

	return cmd
}
