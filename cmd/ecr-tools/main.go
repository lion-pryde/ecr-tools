package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ecr"

	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "get-token":
		getToken(aws.StringValue(flag.Arg(1)))
	default:
		fmt.Println("Command not found.")
		fmt.Println("TODO:... ")
	}
}

func getToken(registryID string) {
	cfg, err := external.LoadDefaultAWSConfig()

	if err != nil {
		panic("Unable to load SDK config, " + err.Error())
	}

	svc := ecr.New(cfg)

	input := &ecr.GetAuthorizationTokenInput{
		RegistryIds: []string{registryID},
	}

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

	for _, token := range result.AuthorizationData {
		token.ProxyEndpoint()
		fmt.Println(aws.StringValue(token.AuthorizationToken))
	}
}
