package messages

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
)

func newReactCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "react <channel> <timestamp> <emoji>",
		Short: "Add a reaction to a message",
		Args:  cobra.ExactArgs(3),
		RunE:  runReact,
	}
}

func runReact(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	// Remove colons if present
	emoji := strings.Trim(args[2], ":")

	if err := c.AddReaction(args[0], args[1], emoji); err != nil {
		return err
	}

	fmt.Printf("Added :%s: reaction\n", emoji)
	return nil
}
