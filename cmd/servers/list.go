package servers

import (
	"fmt"

	"github.com/idsulik/swama/v2/cmd/config"
	"github.com/idsulik/swama/v2/internal/swagger"
	"github.com/spf13/cobra"
)

// newListCommand creates the "servers list" subcommand
func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all servers from a Swagger file",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(cmd.Context(), config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			servers := swagger.NewServers(doc)

			return servers.ListServers()
		},
	}

	return cmd
}
