package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	var (
		owner          = flag.String("o", "owner", "owner name")
		repository     = flag.String("r", "repository", "repositoryName")
		repositoryName = flag.String("repos", "example/web", "repositoryName")
	)
	flag.Parse()
	commitids, err := getAllCommitId(*owner, *repository, 800)
	if err != nil {
		panic(err)
	}
	for _, commit := range commitids {
		fmt.Println(commit)
	}

	imageids, err := getAllImageId(*repositoryName)
	if err != nil {
		panic(err)
	}
	for _, img := range imageids {
		if img.ImageTag == nil {
			continue
		}
		fmt.Println(img)
		fmt.Println(isExists(commitids, *img.ImageTag))
	}
}

func isExists(commitids []string, commitid string) bool {
	for _, id := range commitids {
		if id == commitid {
			return true
		}
	}
	return false
}

func getAllCommitId(owner, repo string, max int) ([]string, error) {
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
			if len(commitids) >= max {
				break
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return commitids, nil
}

func getAllImageId(repositoryName string) ([]*ecr.ImageIdentifier, error) {
	imageIds := []*ecr.ImageIdentifier{}

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	config := aws.NewConfig().WithRegion("ap-northeast-1")
	svc := ecr.New(sess, config)

	params := &ecr.ListImagesInput{
		RepositoryName: aws.String(repositoryName),
	}

	for {
		resp, err := svc.ListImages(params)
		if err != nil {
			return nil, err
		}
		imageIds = append(imageIds, resp.ImageIds...)
		if resp.NextToken == nil {
			break
		}
		params.NextToken = aws.String(*resp.NextToken)
	}
	return imageIds, nil
}
