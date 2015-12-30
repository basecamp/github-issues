package cmd

import (
	"github.com/spf13/cobra"

	"github.com/highrisehq/github-issues/config"
)

// issuesCmd respresents the issues command
var issuesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize github-issues CLI tool and generate config file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a Cli library for Go that empowers applications. This
application is a tool to generate the needed files to quickly create a Cobra
application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.GenerateConfig()
	},
}

func init() {
	RootCmd.AddCommand(issuesCmd)

	// Here you will define your flags and configuration settings

	// Cobra supports Persistent Flags which will work for this command and all subcommands
	// issuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly
	// issuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle" )

}
