package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"swama/pkg/converter"
	"swama/pkg/openapi"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert an endpoint to curl or fetch",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		parts := strings.Split(args[0], " ")
		var method, endpoint string
		if len(parts) == 1 {
			method = "GET"
			endpoint = parts[0]
		} else {
			method = parts[0]
			endpoint = strings.Join(parts[1:], " ")
		}
		convertType := args[1]

		swagger, err := openapi.ParseSwaggerFile(ctx, config.SwaggerPath)
		if err != nil {
			return fmt.Errorf("error parsing Swagger file: %w", err)
		}

		converter, err := converter.NewConverter(convertType)
		if err != nil {
			return fmt.Errorf("error creating converter: %w", err)
		}

		value, err := converter.ConvertEndpoint(swagger, method, endpoint)

		if err != nil {
			return fmt.Errorf("error converting endpoint: %w", err)
		}

		fmt.Println(value)
		return nil
	},
}
