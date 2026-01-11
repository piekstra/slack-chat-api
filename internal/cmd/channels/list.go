package channels

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
	"github.com/piekstra/slack-cli/internal/output"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all channels",
		RunE:  runList,
	}

	cmd.Flags().String("types", "", "Channel types (public_channel,private_channel,mpim,im)")
	cmd.Flags().Bool("exclude-archived", true, "Exclude archived channels")
	cmd.Flags().Int("limit", 100, "Maximum channels to return")

	return cmd
}

func runList(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	types, _ := cmd.Flags().GetString("types")
	excludeArchived, _ := cmd.Flags().GetBool("exclude-archived")
	limit, _ := cmd.Flags().GetInt("limit")

	channels, err := c.ListChannels(types, excludeArchived, limit)
	if err != nil {
		return err
	}

	if output.JSON {
		data, _ := json.MarshalIndent(channels, "", "  ")
		fmt.Println(string(data))
		return nil
	}

	if len(channels) == 0 {
		fmt.Println("No channels found")
		return nil
	}

	fmt.Printf("%-12s %-30s %s\n", "ID", "NAME", "MEMBERS")
	fmt.Println(strings.Repeat("-", 60))
	for _, ch := range channels {
		private := ""
		if ch.IsPrivate {
			private = " (private)"
		}
		fmt.Printf("%-12s %-30s %d%s\n", ch.ID, ch.Name, ch.NumMembers, private)
	}

	return nil
}
