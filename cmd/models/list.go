package models

import (
	"fmt"

	"github.com/idsulik/swama/cmd/config"
	"github.com/idsulik/swama/internal/swagger"
	"github.com/spf13/cobra"
)

// newListCommand creates the "models list" subcommand
func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all API models from a Swagger file",
		Example: "swama models list",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			models := swagger.NewModels(doc)

			return models.ListModels()
		},
	}

	return cmd
}
