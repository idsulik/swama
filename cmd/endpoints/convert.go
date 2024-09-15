package endpoints

import (
	"fmt"

	"github.com/idsulik/swama/cmd/config"
	"github.com/idsulik/swama/internal/swagger"
	"github.com/spf13/cobra"
)

type convertConfig struct {
	method   string
	endpoint string
	toType   string
}

// Command-specific flags for the convert command
var convertCfg = convertConfig{}

// newConvertCommand creates the "endpoints convert" subcommand
func newConvertCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert an endpoint to curl or fetch",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			endpoints := swagger.NewEndpoints(doc)

			return endpoints.ConvertEndpoint(convertCfg.method, convertCfg.endpoint, convertCfg.toType)
		},
	}

	cmd.Flags().StringVarP(&convertCfg.method, "method", "m", "", "Method of the endpoint to convert")
	cmd.Flags().StringVarP(&convertCfg.endpoint, "endpoint", "e", "", "Endpoint to convert")
	cmd.Flags().StringVarP(&convertCfg.toType, "type", "t", "", "Type to convert to (curl, fetch)")

	return cmd
}
