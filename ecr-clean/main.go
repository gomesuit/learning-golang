package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
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
	//for _, commit := range commitids {
	//	fmt.Println(commit)
	//}

	imageids, err := getAllImageId(*repositoryName)
	if err != nil {
		panic(err)
	}

	deleteImgs := getDeleteImgs(imageids, commitids)
	//fmt.Println(deleteImgs)

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	config := aws.NewConfig().WithRegion("ap-northeast-1")
	svc := ecr.New(sess, config)

	i := 0
	for i = 0; i < int(len(deleteImgs)/100); i++ {
		err := deleteImages(svc, *repositoryName, deleteImgs[i*100:(i+1)*100])

		if err != nil {
			fmt.Errorf("deleting images in repo %v: %v", *repositoryName, err)
			return
		}
	}

	err = deleteImages(svc, *repositoryName, deleteImgs[i*100:])

	if err != nil {
		fmt.Errorf("deleting images in repo %v: %v", *repositoryName, err)
		return
	}

	log.Printf("deleted %v images in repo %v", len(deleteImgs), *repositoryName)
}

func deleteImages(ecrCli *ecr.ECR, repoName string, images []*ecr.ImageIdentifier) error {
	_, err := ecrCli.BatchDeleteImage(&ecr.BatchDeleteImageInput{
		RepositoryName: aws.String(repoName),
		ImageIds:       images,
	})
	if err != nil {
		return fmt.Errorf("deleting images in repo %v: %v", repoName, err)
	}

	return nil
}

func getDeleteImgs(imgs []*ecr.ImageIdentifier, commitids []string) []*ecr.ImageIdentifier {
	deleteImgs := []*ecr.ImageIdentifier{}
	for _, img := range imgs {
		if img.ImageTag == nil {
			deleteImgs = append(deleteImgs, img)
			continue
		}
		if !isExists(commitids, *img.ImageTag) {
			deleteImgs = append(deleteImgs, img)
		}
	}
	return deleteImgs
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
