package projects

import (
	"os"
	"regexp"

	git "github.com/libgit2/git2go"
)

type Project struct {
	Directory string
}

func Current() *Project {
	directory, _ := os.Getwd()

	return &Project{Directory: directory}
}

func (project *Project) GithubOwner() string {
	originUrl := gitOriginURL(project)
	re := regexp.MustCompile("github.com:([[:word:]]+)/([[:word:]]+).git")
	matches := re.FindStringSubmatch(originUrl)
	if len(matches) == 3 {
		return matches[1]
	} else {
		return ""
	}
}

func (project *Project) GithubRepo() string {
	originUrl := gitOriginURL(project)
	re := regexp.MustCompile("github.com:([[:word:]]+)/([[:word:]]+).git")
	matches := re.FindStringSubmatch(originUrl)
	if len(matches) == 3 {
		return matches[2]
	} else {
		return ""
	}
}

func gitOriginURL(project *Project) string {
	repo, _ := git.OpenRepository(project.Directory)
	config, _ := repo.Config()
	url, _ := config.LookupString("remote.origin.url")
	return url
}
