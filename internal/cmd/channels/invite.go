package channels

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
)

func newInviteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "invite <channel-id> <user-id>...",
		Short: "Invite users to a channel",
		Args:  cobra.MinimumNArgs(2),
		RunE:  runInvite,
	}
}

func runInvite(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	if err := c.InviteToChannel(args[0], args[1:]); err != nil {
		return err
	}

	fmt.Printf("Invited %d user(s) to channel %s\n", len(args)-1, args[0])
	return nil
}
