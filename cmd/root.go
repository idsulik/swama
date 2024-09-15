package cmd

import (
	"context"

	"github.com/idsulik/swama/pkg/openapi"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "swama",
	Short: "CLI tool for Swagger/OpenAPI operations",
	Long:  `A simple CLI tool to list, view and convert Swagger endpoints.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if config.SwaggerPath == "" {
			config.SwaggerPath = openapi.LocateSwaggerFile()
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
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(convertCmd)
	rootCmd.SetHelpCommand(nil)
}
