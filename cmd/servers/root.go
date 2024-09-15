package servers

import (
	"github.com/spf13/cobra"
)

func NewServersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "servers",
		Aliases: []string{"server"},
		Short:   "Interact with servers",
	}

	cmd.AddCommand(newListCommand())

	return cmd
}
