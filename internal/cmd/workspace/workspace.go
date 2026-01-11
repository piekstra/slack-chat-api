package workspace

import (
	"github.com/spf13/cobra"
)

// NewCmd creates the workspace command with all subcommands
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workspace",
		Aliases: []string{"ws", "team"},
		Short:   "Get workspace information",
	}

	cmd.AddCommand(newInfoCmd())

	return cmd
}
