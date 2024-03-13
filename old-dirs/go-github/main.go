package main

import (
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	var (
		owner      = flag.String("o", "owner", "owner name")
		repository = flag.String("r", "repository", "repositoryName")
	)
	flag.Parse()
	commitids, err := getAllCommitId(*owner, *repository)
	if err != nil {
		panic(err)
	}
	for _, commit := range commitids {
		fmt.Println(commit)
	}
}

func getAllCommitId(owner, repo string) ([]string, error) {
	token := os.Getenv("GITHUB_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	opt := &github.CommitsListOptions{}

	commitids := []string{}
	for {
		commits, resp, err := client.Repositories.ListCommits(owner, repo, opt)
		if err != nil {
			return nil, err
		}
		for _, commit := range commits {
			commitids = append(commitids, *commit.SHA)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return commitids, nil
}
