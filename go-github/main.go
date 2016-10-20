package main

import (
	"fmt"
	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)

	commits, _, err := client.Repositories.ListCommits("gomesuit", "larning-golang", nil)
	if err != nil {
		panic(err)
	}
	for _, commit := range commits {
		fmt.Println(*commit.SHA)
		fmt.Println(commit.Commit.Author.Date)
	}
}
