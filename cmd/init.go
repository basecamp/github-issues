package cmd

import (
	"github.com/spf13/cobra"

	"github.com/highrisehq/github-issues/config"
)

// issuesCmd respresents the issues command
var issuesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize github-issues CLI tool and generate config file",
	Long: `github-issues requires authentication to Github for private repositories.
Go to your Github Profile and get a Personal Token to use. Then, run:

github-issues init -t <your token>`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		config.GenerateConfig(token)
	},
}

func init() {
	RootCmd.AddCommand(issuesCmd)
	issuesCmd.Flags().StringP("token", "t", "", "Github auth token to initialize config with")
}
