package channels

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
)

func newSetPurposeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set-purpose <channel-id> <purpose>",
		Short: "Set channel purpose",
		Args:  cobra.ExactArgs(2),
		RunE:  runSetPurpose,
	}
}

func runSetPurpose(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	if err := c.SetChannelPurpose(args[0], args[1]); err != nil {
		return err
	}

	fmt.Printf("Set purpose for channel %s\n", args[0])
	return nil
}
