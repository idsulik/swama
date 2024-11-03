package mockserver

import (
	"fmt"

	"github.com/idsulik/swama/v2/cmd/config"
	mockserver2 "github.com/idsulik/swama/v2/internal/mockserver"
	"github.com/idsulik/swama/v2/internal/swagger"
	"github.com/spf13/cobra"
)

type runConfig struct {
	host                string
	port                int
	delay               int
	defaultResponseCode string
	defaultResponseType string
}

// Command-specific flags for the run command
var runCfg = runConfig{}

// newRunCommand creates the "tags run" subcommand
func newRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run details of a specific tag",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := swagger.LoadSwaggerFile(cmd.Context(), config.SwaggerPath)

			if err != nil {
				return fmt.Errorf("failed to load Swagger file: %w", err)
			}

			mockserver := mockserver2.NewMockServer(doc)

			return mockserver.Run(
				mockserver2.RunOptions{
					Host:                runCfg.host,
					Port:                runCfg.port,
					Delay:               runCfg.delay,
					DefaultResponseCode: runCfg.defaultResponseCode,
					DefaultResponseType: runCfg.defaultResponseType,
				},
			)
		},
	}

	cmd.Flags().StringVarP(&runCfg.host, "host", "", "127.0.0.1", "Host to run the mock server on")
	cmd.Flags().IntVarP(&runCfg.port, "port", "p", 8080, "Port to run the mock server on")
	cmd.Flags().IntVarP(&runCfg.delay, "delay", "d", 0, "Delay in milliseconds to simulate network latency")
	cmd.Flags().StringVarP(
		&runCfg.defaultResponseCode,
		"default-response-code",
		"",
		"200",
		"Default response code to use",
	)
	cmd.Flags().StringVarP(
		&runCfg.defaultResponseType,
		"default-response-type",
		"",
		"json",
		"Default response type to use",
	)

	return cmd
}
