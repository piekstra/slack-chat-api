package messages

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
)

func newDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <channel> <timestamp>",
		Short: "Delete a message",
		Args:  cobra.ExactArgs(2),
		RunE:  runDelete,
	}
}

func runDelete(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	if err := c.DeleteMessage(args[0], args[1]); err != nil {
		return err
	}

	fmt.Println("Message deleted")
	return nil
}
