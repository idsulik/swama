package tags

import (
	"fmt"

	"github.com/idsulik/swama/v2/cmd/config"
	"github.com/idsulik/swama/v2/internal/swagger"
	"github.com/spf13/cobra"
)

// newListCommand creates the "tags list" subcommand
func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all tags from a Swagger file",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(cmd.Context(), config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			tags := swagger.NewTags(doc)

			return tags.ListTags()
		},
	}

	return cmd
}
