package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

// Config ...

func main() {

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	svc := ecr.New(cfg)
	input := &ecr.GetAuthorizationTokenInput{}

	result, err := svc.GetAuthorizationTokenRequest(input).Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result.String())
}
