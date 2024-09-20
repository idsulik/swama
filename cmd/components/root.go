package components

import (
	"github.com/spf13/cobra"
)

func NewComponentsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "components",
		Aliases: []string{"component", "compo", "schemas"},
		Short:   "Interact with API components",
	}

	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newViewCommand())

	return cmd
}
