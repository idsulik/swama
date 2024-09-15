package tags

import (
	"github.com/spf13/cobra"
)

func NewTagsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tags",
		Aliases: []string{"tag"},
		Short:   "Interact with tags",
	}

	cmd.AddCommand(newListCommand())

	return cmd
}
