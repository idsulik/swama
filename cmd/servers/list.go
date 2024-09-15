package servers

import (
	"fmt"

	"github.com/idsulik/swama/cmd/config"
	"github.com/idsulik/swama/internal/swagger"
	"github.com/spf13/cobra"
)

// newListCommand creates the "servers list" subcommand
func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all servers from a Swagger file",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			servers := swagger.NewServers(doc)

			return servers.ListServers()
		},
	}

	return cmd
}
