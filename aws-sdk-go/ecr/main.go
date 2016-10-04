package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	config := aws.NewConfig().WithRegion("ap-northeast-1")
	svc := ecr.New(sess, config)
	params := &ecr.DescribeRepositoriesInput{
		MaxResults: aws.Int64(100),
	}
	resp, err := svc.DescribeRepositories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
