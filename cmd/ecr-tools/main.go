package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/lyon-pryde/ecr-tools/pkg/collector"
	"plugin"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)
var AWSAccessID      = os.Getenv("AWS_ACCESS_KEY_ID")
var AWSDefaultRegion      = os.Getenv("AWS_DEFAULT_REGION")
var AWSSecretAccessKey      = os.Getenv("AWS_SECRET_ACCESS_KEY")

var AWSAccessIDUsage = "AWS Access ID overrides env variable AWS_ACCESS_KEY_ID"
var AWSSecretAccessKeyUsage = "AWS Access ID overrides env variable AWS_SECRET_ACCESS_KEY"
var AWSDefaultRegionUsage = "AWS Default Region overrides env variable AWS_DEFAULT_REGION"

var configFlags = flag.NewFlagSet("get-token", flag.ExitOnError)

var awsAccessID = configFlags.String("aws-access-id", AWSAccessID, AWSAccessIDUsage)
var awsSecretAccessKey = configFlags.String("aws-secret-access-key", AWSSecretAccessKey, AWSSecretAccessKeyUsage)
var awsDefaultRegion = configFlags.String("aws-default-region", AWSDefaultRegion, AWSDefaultRegionUsage)

func init() {
	configFlags.StringVar(awsAccessID, "id", "", "Description")
	configFlags.StringVar(awsSecretAccessKey, "key", "", "Description")
	configFlags.StringVar(awsDefaultRegion, "region", "", "Description")
}

func main() {
	configFlags.Parse(os.Args[2:])

	var AccessID string
	AccessID = *awsAccessID
	var SecretAccessKey string
	SecretAccessKey = *awsSecretAccessKey
	var DefaultRegion string
	DefaultRegion = *awsDefaultRegion

	credentialsProvider := Provider{
		awsAccessID:        AccessID,
		awsSecretAccessKey: SecretAccessKey,
		awsDefaultRegion:   DefaultRegion,
	}
	fmt.Println(credentialsProvider)
	switch os.Args[1] {
		case "get-token":
			valid, err := credentialsProvider.AuthValid()

			getToken(credentialsProvider)
			if valid {
				if err != nil {
					fmt.Println("fuck")
				}else{
					fmt.Println("Invalid Auth")
					flag.PrintDefaults()
					os.Exit(67)
				}
			} else {
				fmt.Println(err)
				flag.PrintDefaults()
				os.Exit(66)
			}
		default:
			flag.PrintDefaults()
	}
}

func getToken(p Provider) {
	// Open plugin, and load it into the process.
	//p, err := plugin.Open("./plugin/plugin.so")
	//if err != nil {
	//	exitErrorf("failed to open plugin, %s, %v", pluginFilename, err)
	//}
	// Create a new Credentials value which will source the provider's Retrieve
	// and IsExpired functions from the plugin.
	//creds, err := plugincreds.NewCredentials(p)
	creds := aws.Credentials{
		AccessKeyID:
	}
	// Load the config and set with the newly created credentials that
	// will be sourced using the plugin's functionality.
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		exitErrorf("failed to load config, %v", err)
	}
	cfg.Credentials = creds



	fmt.Println("Config", cfg)
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
	fmt.Println(result)
	for _, token := range result.AuthorizationData {
		fmt.Println(token.AuthorizationToken)
	}
}
