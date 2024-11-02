package mockserver

import (
	"fmt"

	"github.com/idsulik/swama/cmd/config"
	mockserver2 "github.com/idsulik/swama/internal/mockserver"
	"github.com/idsulik/swama/internal/swagger"
	"github.com/spf13/cobra"
)

type runConfig struct {
	host  string
	port  int
	delay int
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
					Host:  runCfg.host,
					Port:  runCfg.port,
					Delay: runCfg.delay,
				},
			)
		},
	}

	cmd.Flags().StringVarP(&runCfg.host, "host", "", "127.0.0.1", "Host to run the mock server on")
	cmd.Flags().IntVarP(&runCfg.port, "port", "p", 8080, "Port to run the mock server on")
	cmd.Flags().IntVarP(&runCfg.delay, "delay", "d", 0, "Delay in milliseconds to simulate network latency")

	return cmd
}
