package channels

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
)

func newArchiveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "archive <channel-id>",
		Short: "Archive a channel",
		Args:  cobra.ExactArgs(1),
		RunE:  runArchive,
	}
}

func runArchive(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	if err := c.ArchiveChannel(args[0]); err != nil {
		return err
	}

	fmt.Printf("Archived channel: %s\n", args[0])
	return nil
}
