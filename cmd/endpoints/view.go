package endpoints

import (
	"fmt"

	"github.com/idsulik/swama/cmd/config"
	"github.com/idsulik/swama/internal/swagger"
	"github.com/spf13/cobra"
)

type viewConfig struct {
	method   string
	endpoint string
}

// Command-specific flags for the view command
var viewCfg = viewConfig{}

// newViewCommand creates the "endpoints view" subcommand
func newViewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "View details of a specific endpoint",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(cmd.Context(), config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			endpoints := swagger.NewEndpoints(doc)

			return endpoints.ViewEndpoint(
				swagger.ViewOptions{
					Method:   viewCfg.method,
					Endpoint: viewCfg.endpoint,
				},
			)
		},
	}

	cmd.Flags().StringVarP(&viewCfg.method, "method", "m", "", "Method of the endpoint to view")
	cmd.Flags().StringVarP(&viewCfg.endpoint, "endpoint", "e", "", "Endpoint to view")

	return cmd
}
