package messages

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
	"github.com/piekstra/slack-cli/internal/output"
)

func newHistoryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "history <channel>",
		Short: "Get channel message history",
		Args:  cobra.ExactArgs(1),
		RunE:  runHistory,
	}

	cmd.Flags().Int("limit", 20, "Maximum messages to return")
	cmd.Flags().String("oldest", "", "Only messages after this timestamp")
	cmd.Flags().String("latest", "", "Only messages before this timestamp")

	return cmd
}

func runHistory(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	limit, _ := cmd.Flags().GetInt("limit")
	oldest, _ := cmd.Flags().GetString("oldest")
	latest, _ := cmd.Flags().GetString("latest")

	messages, err := c.GetChannelHistory(args[0], limit, oldest, latest)
	if err != nil {
		return err
	}

	if output.JSON {
		data, _ := json.MarshalIndent(messages, "", "  ")
		fmt.Println(string(data))
		return nil
	}

	if len(messages) == 0 {
		fmt.Println("No messages found")
		return nil
	}

	for _, m := range messages {
		ts := formatTimestamp(m.TS)
		text := truncate(m.Text, 80)
		fmt.Printf("[%s] %s: %s\n", ts, m.User, text)
	}

	return nil
}
