package cmd

import (
	"fmt"

	"github.com/idsulik/swama/pkg/openapi"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View details of a specific endpoint",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := args[0]

		ctx := cmd.Context()
		swagger, err := openapi.ParseSwaggerFile(ctx, config.SwaggerPath)
		if err != nil {
			return fmt.Errorf("error parsing Swagger file: %w", err)
		}

		details, err := openapi.ViewEndpointDetails(swagger, endpoint)

		if err != nil {
			return err
		}

		fmt.Println(details)
		return nil
	},
}
