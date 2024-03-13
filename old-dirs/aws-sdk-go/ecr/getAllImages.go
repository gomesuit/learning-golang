package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func main() {
	var (
		repositoryName = flag.String("repos", "example/web", "repositoryName")
	)
	flag.Parse()
	imageids, err := getAllImageId(*repositoryName)
	if err != nil {
		panic(err)
	}
	for _, img := range imageids {
		fmt.Println(img)
	}
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
