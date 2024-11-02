package cmd

import (
	"context"

	"github.com/idsulik/swama/v2/cmd/components"
	"github.com/idsulik/swama/v2/cmd/config"
	"github.com/idsulik/swama/v2/cmd/endpoints"
	"github.com/idsulik/swama/v2/cmd/info"
	"github.com/idsulik/swama/v2/cmd/mockserver"
	"github.com/idsulik/swama/v2/cmd/servers"
	"github.com/idsulik/swama/v2/cmd/tags"
	"github.com/idsulik/swama/v2/internal/swagger"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "swama",
	Short: "Swama is a CLI tool for interacting with Swagger/OpenAPI definitions",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if config.SwaggerPath == "" {
			config.SwaggerPath = swagger.LocateSwaggerFile()
		}
	},
}

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&config.SwaggerPath,
		"file",
		"f",
		"",
		"Path to the Swagger JSON/YAML file. If not provided, the tool will try to locate it.",
	)

	rootCmd.AddCommand(info.NewInfoCommand())
	rootCmd.AddCommand(endpoints.NewEndpointsCommand())
	rootCmd.AddCommand(components.NewComponentsCommand())
	rootCmd.AddCommand(servers.NewServersCommand())
	rootCmd.AddCommand(tags.NewTagsCommand())
	rootCmd.AddCommand(mockserver.NewMockCommand())
}
