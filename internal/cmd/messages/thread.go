package messages

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
	"github.com/piekstra/slack-cli/internal/output"
)

func newThreadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "thread <channel> <thread-ts>",
		Short: "Get thread replies",
		Args:  cobra.ExactArgs(2),
		RunE:  runThread,
	}

	cmd.Flags().Int("limit", 100, "Maximum replies to return")

	return cmd
}

func runThread(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	limit, _ := cmd.Flags().GetInt("limit")

	messages, err := c.GetThreadReplies(args[0], args[1], limit)
	if err != nil {
		return err
	}

	if output.JSON {
		data, _ := json.MarshalIndent(messages, "", "  ")
		fmt.Println(string(data))
		return nil
	}

	if len(messages) == 0 {
		fmt.Println("No replies found")
		return nil
	}

	for _, m := range messages {
		ts := formatTimestamp(m.TS)
		text := truncate(m.Text, 80)
		fmt.Printf("[%s] %s: %s\n", ts, m.User, text)
	}

	return nil
}
