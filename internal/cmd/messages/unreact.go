package messages

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
)

func newUnreactCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "unreact <channel> <timestamp> <emoji>",
		Short: "Remove a reaction from a message",
		Args:  cobra.ExactArgs(3),
		RunE:  runUnreact,
	}
}

func runUnreact(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	// Remove colons if present
	emoji := strings.Trim(args[2], ":")

	if err := c.RemoveReaction(args[0], args[1], emoji); err != nil {
		return err
	}

	fmt.Printf("Removed :%s: reaction\n", emoji)
	return nil
}
