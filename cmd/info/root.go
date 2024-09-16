package info

import (
	"github.com/spf13/cobra"
)

func NewInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Interact with info",
	}

	cmd.AddCommand(vewViewCommand())

	return cmd
}
