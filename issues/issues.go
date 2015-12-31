package issues

import (
	"fmt"
	"sync"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/highrisehq/github-issues/config"
	"github.com/highrisehq/github-issues/logger"
)

type Query struct {
	Labels []string
	State  string
	Owner  string
	Repo   string
}

func (query *Query) Execute() {
	client := client()

	options := github.IssueListByRepoOptions{State: query.State, Labels: query.Labels}
	issues, _, err := client.Issues.ListByRepo(query.Owner, query.Repo, &options)
	if err != nil {
		logger.Error(err.Error())
	} else if len(issues) == 0 {
		logger.Info("no pulls")
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
