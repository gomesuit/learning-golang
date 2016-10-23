package main

import (
	"fmt"
	"github.com/google/go-github/github"
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
	client := github.NewClient(nil)
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
