package endpoints

import (
	"fmt"

	"github.com/idsulik/swama/cmd/config"
	"github.com/idsulik/swama/internal/swagger"
	"github.com/spf13/cobra"
)

type listConfig struct {
	method   string
	endpoint string
	tag      string
	group    string
}

// Command-specific flags for the list command
var listCfg = listConfig{}

// newListCommand creates the "endpoints list" subcommand
func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists all API endpoints from a Swagger file",
		Example: "swama endpoints list --method GET --tag user",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			endpoints := swagger.NewEndpoints(doc)

			return endpoints.ListEndpoints(listCfg.method, listCfg.endpoint, listCfg.tag, listCfg.group)
		},
	}

	cmd.Flags().StringVarP(&listCfg.method, "method", "m", "", "Filter by method")
	cmd.Flags().StringVarP(&listCfg.endpoint, "endpoint", "e", "", "Filter by endpoint, supports wildcard")
	cmd.Flags().StringVarP(&listCfg.tag, "tag", "t", "", "Filter by tag")
	cmd.Flags().StringVarP(&listCfg.group, "group", "g", "tag", "Group output by tag, method")

	return cmd
}
