package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	commitids, err := getAllCommitId("gomesuit", "jansible")
	if err != nil {
		panic(err)
	}
	for _, commit := range commitids {
		fmt.Println(commit)
	}
}

func getAllCommitId(owner, repo string) ([]string, error) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "xxxxxxxxxxxxxxxxxxxxxxxxxx"},
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
