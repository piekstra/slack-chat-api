package channels

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
)

func newSetTopicCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set-topic <channel-id> <topic>",
		Short: "Set channel topic",
		Args:  cobra.ExactArgs(2),
		RunE:  runSetTopic,
	}
}

func runSetTopic(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	if err := c.SetChannelTopic(args[0], args[1]); err != nil {
		return err
	}

	fmt.Printf("Set topic for channel %s\n", args[0])
	return nil
}
