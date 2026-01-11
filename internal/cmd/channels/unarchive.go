package channels

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
)

func newUnarchiveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "unarchive <channel-id>",
		Short: "Unarchive a channel",
		Args:  cobra.ExactArgs(1),
		RunE:  runUnarchive,
	}
}

func runUnarchive(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	if err := c.UnarchiveChannel(args[0]); err != nil {
		return err
	}

	fmt.Printf("Unarchived channel: %s\n", args[0])
	return nil
}
