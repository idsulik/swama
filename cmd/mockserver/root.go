package mockserver

import (
	"github.com/spf13/cobra"
)

func NewMockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mock-server",
		Aliases: []string{"mock-server"},
		Short:   "Interact with mock server",
	}

	cmd.AddCommand(newRunCommand())

	return cmd
}
