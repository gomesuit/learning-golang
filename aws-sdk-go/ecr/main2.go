package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func main() {
	repositoryName := "medical/web"

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	config := aws.NewConfig().WithRegion("ap-northeast-1")
	svc := ecr.New(sess, config)

	params := &ecr.ListImagesInput{
		RepositoryName: aws.String(repositoryName),
		Filter: &ecr.ListImagesFilter{
			TagStatus: aws.String("UNTAGGED"),
		},
	}
	resp, err := svc.ListImages(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp.ImageIds)

	if resp.NextToken == nil {
		return
	}

	for {
		params := &ecr.ListImagesInput{
			RepositoryName: aws.String(repositoryName),
			Filter: &ecr.ListImagesFilter{
				TagStatus: aws.String("UNTAGGED"),
			},
			NextToken: aws.String(*resp.NextToken),
		}
		resp, err := svc.ListImages(params)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(resp.ImageIds)
		if resp.NextToken == nil {
			break
		}
	}
}
