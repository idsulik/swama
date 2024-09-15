package endpoints

import (
	"github.com/spf13/cobra"
)

func NewEndpointsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "endpoints",
		Aliases: []string{"endpoint", "ep"},
		Short:   "Interact with API endpoints",
	}

	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newViewCommand())
	cmd.AddCommand(newConvertCommand())

	return cmd
}
