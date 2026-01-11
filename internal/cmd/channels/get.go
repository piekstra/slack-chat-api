package channels

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
	"github.com/piekstra/slack-cli/internal/output"
)

func newGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <channel-id>",
		Short: "Get channel information",
		Args:  cobra.ExactArgs(1),
		RunE:  runGet,
	}
}

func runGet(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	channel, err := c.GetChannelInfo(args[0])
	if err != nil {
		return err
	}

	if output.JSON {
		data, _ := json.MarshalIndent(channel, "", "  ")
		fmt.Println(string(data))
		return nil
	}

	fmt.Printf("ID:       %s\n", channel.ID)
	fmt.Printf("Name:     %s\n", channel.Name)
	fmt.Printf("Private:  %t\n", channel.IsPrivate)
	fmt.Printf("Archived: %t\n", channel.IsArchived)
	fmt.Printf("Members:  %d\n", channel.NumMembers)
	if channel.Topic.Value != "" {
		fmt.Printf("Topic:    %s\n", channel.Topic.Value)
	}
	if channel.Purpose.Value != "" {
		fmt.Printf("Purpose:  %s\n", channel.Purpose.Value)
	}

	return nil
}
