package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/cmd/channels"
	"github.com/piekstra/slack-cli/internal/cmd/config"
	"github.com/piekstra/slack-cli/internal/cmd/messages"
	"github.com/piekstra/slack-cli/internal/cmd/users"
	"github.com/piekstra/slack-cli/internal/cmd/workspace"
	"github.com/piekstra/slack-cli/internal/output"
	"github.com/piekstra/slack-cli/internal/version"
)

var rootCmd = &cobra.Command{
	Use:   "slack-cli",
	Short: "A CLI tool for interacting with Slack",
	Long: `slack-cli is a command-line interface for Slack.

It provides commands for managing channels, users, messages,
and other Slack workspace operations.

Configure your API token with:
  slack-cli config set-token <your-token>

Or set the SLACK_API_TOKEN environment variable.`,
	Version: version.Version,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&output.JSON, "json", false, "Output in JSON format")

	// Set custom version template to include commit and build date
	rootCmd.SetVersionTemplate("slack-cli " + version.Info() + "\n")

	// Add subcommands
	rootCmd.AddCommand(channels.NewCmd())
	rootCmd.AddCommand(users.NewCmd())
	rootCmd.AddCommand(messages.NewCmd())
	rootCmd.AddCommand(workspace.NewCmd())
	rootCmd.AddCommand(config.NewCmd())
}
