package messages

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
	"github.com/piekstra/slack-cli/internal/output"
)

func newSendCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send <channel> <text>",
		Short: "Send a message to a channel",
		Long: `Send a message to a channel.

By default, messages are sent using Slack Block Kit formatting for a more
refined appearance. Use --simple to send plain text messages instead.`,
		Args: cobra.ExactArgs(2),
		RunE: runSend,
	}

	cmd.Flags().String("thread", "", "Thread timestamp for reply")
	cmd.Flags().String("blocks", "", "Block Kit blocks as JSON array (overrides default block formatting)")
	cmd.Flags().Bool("simple", false, "Send as plain text without block formatting")

	return cmd
}

func runSend(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	threadTS, _ := cmd.Flags().GetString("thread")
	blocksJSON, _ := cmd.Flags().GetString("blocks")
	simple, _ := cmd.Flags().GetBool("simple")

	var blocks []interface{}
	if blocksJSON != "" {
		if err := json.Unmarshal([]byte(blocksJSON), &blocks); err != nil {
			return fmt.Errorf("invalid blocks JSON: %w", err)
		}
	} else if !simple {
		// Default to block style for a more refined appearance
		blocks = buildDefaultBlocks(args[1])
	}

	msg, err := c.SendMessage(args[0], args[1], threadTS, blocks)
	if err != nil {
		return err
	}

	if output.JSON {
		data, _ := json.MarshalIndent(msg, "", "  ")
		fmt.Println(string(data))
		return nil
	}

	fmt.Printf("Message sent (ts: %s)\n", msg.TS)
	return nil
}
