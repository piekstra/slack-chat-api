package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/piekstra/slack-cli/internal/keychain"
	"github.com/piekstra/slack-cli/internal/output"
)

type deleteTokenOptions struct {
	force bool
	stdin io.Reader // For testing
}

func newDeleteTokenCmd() *cobra.Command {
	opts := &deleteTokenOptions{}

	cmd := &cobra.Command{
		Use:   "delete-token",
		Short: "Delete the stored Slack API token",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDeleteToken(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Skip confirmation prompt")

	return cmd
}

func runDeleteToken(opts *deleteTokenOptions) error {
	// Check if there's actually a stored token to delete
	if !keychain.HasStoredToken() {
		output.Println("No token stored in keychain/config file to delete.")
		if os.Getenv("SLACK_API_TOKEN") != "" {
			output.Println("Note: Token is set via SLACK_API_TOKEN environment variable.")
		}
		return nil
	}

	// Prompt for confirmation unless --force
	if !opts.force {
		reader := opts.stdin
		if reader == nil {
			reader = os.Stdin
		}

		output.Println("About to delete the stored Slack API token.")
		output.Printf("Are you sure? [y/N]: ")

		scanner := bufio.NewScanner(reader)
		if scanner.Scan() {
			confirm := strings.TrimSpace(strings.ToLower(scanner.Text()))
			if confirm != "y" && confirm != "yes" {
				output.Println("Cancelled.")
				return nil
			}
		}
	}

	if err := keychain.DeleteAPIToken(); err != nil {
		return fmt.Errorf("failed to delete token: %w", err)
	}

	if keychain.IsSecureStorage() {
		output.Println("API token deleted from Keychain")
	} else {
		output.Println("API token deleted from config file")
	}
	return nil
}
