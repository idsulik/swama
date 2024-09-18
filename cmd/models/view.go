package models

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

// newViewCommand creates the "models view" subcommand
func newViewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "View details of a specific model",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			models := swagger.NewModels(doc)

			return models.ViewModel(viewCfg.name)
		},
	}

	cmd.Flags().StringVarP(&viewCfg.name, "name", "n", "", "Name of the model to view")

	return cmd
}
