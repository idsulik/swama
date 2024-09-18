package models

import (
	"github.com/spf13/cobra"
)

func NewModelsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "models",
		Aliases: []string{"model", "mod"},
		Short:   "Interact with API models",
	}

	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newViewCommand())

	return cmd
}
