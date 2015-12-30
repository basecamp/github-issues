package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/highrisehq/github-issues/config"
)

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
	Run: func(cmd *cobra.Command, args []string) {
		client := client()

		labels := []string{"beta"}
		options := github.IssueListByRepoOptions{State: "open", Labels: labels}
		issues, _, err := client.Issues.ListByRepo("highrisehq", "server", &options)
		if err != nil {
			fmt.Println(err)
		} else if len(issues) == 0 {
			fmt.Println("no pulls")
		} else {
			var wg sync.WaitGroup
			for _, issue := range issues {
				if issue.PullRequestLinks != nil {
					wg.Add(1)
					go func() {
						defer wg.Done()
						fetchByIssueNumber(*client, *issue.Number)
					}()
					wg.Wait()
				}
			}
		}
	},
}

func fetchByIssueNumber(client github.Client, number int) {
	pull, _, _ := client.PullRequests.Get("highrisehq", "server", number)
	fmt.Printf("%s\n", *pull.Head.Ref)
}

func client() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GetToken()},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	return client
}

//Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(config.ConfigInit)
	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which if defined here will be global for your application

	RootCmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "config file (default is $HOME/.github-issues.yaml)")

	// Cobra also supports local flags which will only run when this action is called directly
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
