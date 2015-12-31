package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/highrisehq/github-issues/config"
	"github.com/highrisehq/github-issues/issues"
	"github.com/highrisehq/github-issues/logger"
	"github.com/highrisehq/github-issues/projects"
)

var debugFlag bool

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "github-issues",
	Short: "CLI Access to querying Github pulls and issues",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your application. For example:

Cobra is a Cli library for Go that empowers applications. This
application is a tool to generate the needed files to quickly create a Cobra
application.`,
	// Uncomment the following line if your bare application has an action associated with it
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debugFlag {
			logger.CurrentLevel = logger.DebugLevel
		}
		config.ConfigInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(ownerArg) == 0 {
			ownerArg = projects.Current().GithubOwner()
		}
		if len(repoArg) == 0 {
			repoArg = projects.Current().GithubRepo()
		}
		if len(repoArg) == 0 || len(ownerArg) == 0 {
			logger.Info("Couldn't find Owner or Repo to query")
			return
		}

		labels := strings.Split(strings.Trim(labelArg, " "), ",")
		issueQuery := issues.Query{
			Labels: labels,
			State:  stateArg,
			Owner:  ownerArg,
			Repo:   repoArg,
		}
		issueQuery.Execute()
	},
}

//Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(-1)
	}
}

var labelArg string
var stateArg string
var ownerArg string
var repoArg string

func init() {
	RootCmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "config file (default is $HOME/.github-issues.yaml)")

	RootCmd.Flags().StringVarP(&labelArg, "labels", "l", "", "comma delimited labels to search by")
	RootCmd.Flags().StringVarP(&stateArg, "state", "s", "open", "Issue State: [open, closed, all]")
	RootCmd.Flags().StringVarP(&ownerArg, "owner", "o", "", "Repo Owner string (your username or organization)")
	RootCmd.Flags().StringVarP(&repoArg, "repo", "r", "", "Repository")

	RootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Debug messages")
}
