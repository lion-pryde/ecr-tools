package main

import (
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ecr"

	"flag"
	"fmt"
	"os"
)

var awsAccessID, awsSecretAccessKey, awsDefaultRegion string

func init() {
	var (
		AWSAccessID      = os.Getenv("AWS_ACCESS_KEY_ID")
		AWSAccessIDUsage = "AWS Access ID overrides env variable AWS_ACCESS_KEY_ID"
	)
	flag.StringVar(&awsAccessID, "aws-access-id", AWSAccessID, AWSAccessIDUsage)
	flag.StringVar(&awsAccessID, "i", AWSAccessID, AWSAccessIDUsage+" (shorthand)")

	var (
		AWSSecretAccessKey      = os.Getenv("AWS_ACCESS_KEY_ID")
		AWSSecretAccessKeyUsage = "AWS Access ID overrides env variable AWS_ACCESS_KEY_ID"
	)
	flag.StringVar(&awsSecretAccessKey, "aws-secret-access-key", AWSSecretAccessKey, AWSSecretAccessKeyUsage)
	flag.StringVar(&awsSecretAccessKey, "k", AWSSecretAccessKey, AWSSecretAccessKeyUsage+" (shorthand)")

	var (
		AWSDefaultRegion      = os.Getenv("AWS_DEFAULT_REGION")
		AWSDefaultRegionUsage = "AWS Default Region overrides env variable AWS_DEFAULT_REGION"
	)
	flag.StringVar(&awsDefaultRegion, "aws-default-region", AWSDefaultRegion, AWSDefaultRegionUsage)
	flag.StringVar(&awsDefaultRegion, "r", AWSDefaultRegion, AWSDefaultRegionUsage+" (shorthand)")
}

func main() {
	flag.Parse()
	switch os.Args[1] {
	case "get-token":
		getToken()
	default:
		fmt.Println("Command not found.")
	}
}

func getToken() {
	// cfg:
	// 	Config{}
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("Unable to load SDK config, " + err.Error())
	}
	svc := ecr.New(aws.Config{
		Region: cfg.Region
	})
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

	fmt.Println(result.AuthorizationData)
}
