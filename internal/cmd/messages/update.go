package messages

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
)

func newUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <channel> <timestamp> <text>",
		Short: "Update an existing message",
		Long: `Update an existing message.

By default, messages are updated using Slack Block Kit formatting for a more
refined appearance. Use --simple to update with plain text instead.`,
		Args: cobra.ExactArgs(3),
		RunE: runUpdate,
	}

	cmd.Flags().String("blocks", "", "Block Kit blocks as JSON array (overrides default block formatting)")
	cmd.Flags().Bool("simple", false, "Update as plain text without block formatting")

	return cmd
}

func runUpdate(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	blocksJSON, _ := cmd.Flags().GetString("blocks")
	simple, _ := cmd.Flags().GetBool("simple")

	var blocks []interface{}
	if blocksJSON != "" {
		if err := json.Unmarshal([]byte(blocksJSON), &blocks); err != nil {
			return fmt.Errorf("invalid blocks JSON: %w", err)
		}
	} else if !simple {
		// Default to block style for a more refined appearance
		blocks = buildDefaultBlocks(args[2])
	}

	if err := c.UpdateMessage(args[0], args[1], args[2], blocks); err != nil {
		return err
	}

	fmt.Println("Message updated")
	return nil
}
