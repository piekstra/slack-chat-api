package channels

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/client"
	"github.com/piekstra/slack-cli/internal/output"
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a new channel",
		Args:  cobra.ExactArgs(1),
		RunE:  runCreate,
	}

	cmd.Flags().Bool("private", false, "Create as private channel")

	return cmd
}

func runCreate(cmd *cobra.Command, args []string) error {
	c, err := client.New()
	if err != nil {
		return err
	}

	isPrivate, _ := cmd.Flags().GetBool("private")

	channel, err := c.CreateChannel(args[0], isPrivate)
	if err != nil {
		return err
	}

	if output.JSON {
		data, _ := json.MarshalIndent(channel, "", "  ")
		fmt.Println(string(data))
		return nil
	}

	fmt.Printf("Created channel: %s (%s)\n", channel.Name, channel.ID)
	return nil
}
