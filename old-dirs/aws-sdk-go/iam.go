package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"reflect"
)

func GetAllAccessKeys(svc *iam.IAM, user string) (*iam.ListAccessKeysOutput, error) {
	params := &iam.ListAccessKeysInput{
		UserName: aws.String(user),
	}
	resp, err := svc.ListAccessKeys(params)
	return resp, err
}

func main() {
	config := aws.NewConfig().WithRegion("ap-northeast-1")
	sess, err := session.NewSession(config)
	svc := iam.New(sess)

	fmt.Println(reflect.ValueOf(svc).Type())

	params := &iam.ListUsersInput{}
	resp, err := svc.ListUsers(params)

	for _, user := range resp.Users {
		fmt.Println(aws.StringValue(user.UserName))
		resp, err := GetAllAccessKeys(svc, aws.StringValue(user.UserName))
		fmt.Println(resp)

		for _, key := range resp.AccessKeyMetadata {
			params := &iam.GetAccessKeyLastUsedInput{
				AccessKeyId: key.AccessKeyId,
			}
			aaa, err := svc.GetAccessKeyLastUsed(params)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(aaa)
		}

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// print json
	//fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
}
