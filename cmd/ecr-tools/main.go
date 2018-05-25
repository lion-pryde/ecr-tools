package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ecr"

	"flag"
	"fmt"
	"os"
)

const (
	awsAccessKeyIDEnvVar = "AWS_ACCESS_KEY_ID"
	awsAccessKeyEnvVar   = "AWS_ACCESS_KEY"

	awsSecreteAccessKeyEnvVar = "AWS_SECRET_ACCESS_KEY"
	awsSecreteKeyEnvVar       = "AWS_SECRET_KEY"

	awsECRegistyID = "AWS_ECR_REGISTRY_ID"
	awsECRegisty   = "AWS_ECR_REGISTRY"
)

var (
	registryID    = os.Getenv("AWS_REGISTRY_ID")
	defaultRegion = os.Getenv("AWS_REGISTRY_ID")
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "get-token":

		ec := NewECR()
		ec.SetRegistyID(registryID)
		ec.SetDefaultRegion(defaultRegion)

		return

	default:
		fmt.Println("Command not found: ")
		fmt.Println(fmt.Sprintf(`Commands:
			get-token
		Inputs as ENV Variables:
		
		export %s<Value> or export  %s=<Value>
		
		export %s=<Value> or export %s=<Value>
	
		export %s=<Value> or export %s=<Value>
		`, awsAccessKeyIDEnvVar, awsAccessKeyEnvVar, awsSecreteAccessKeyEnvVar, awsSecreteKeyEnvVar, awsECRegistyID, awsECRegisty))
	}
}

// ECR -
type ECR struct {
	ecr.ECR
	registryID    string
	defaultRegion string
}

// NewECR -
func NewECR() *ECR {
	return &ECR{}
}

// GetToken - GetToken
func (cr ECR) GetToken() (token string, erroa error) {
	cfg, err := external.LoadDefaultAWSConfig()

	svc := ecr.New(cfg)

	var input *ecr.GetAuthorizationTokenInput

	if len(cr.registryID) > 0 {
		input = &ecr.GetAuthorizationTokenInput{
			RegistryIds: []string{cr.registryID},
		}
	} else {
		input = &ecr.GetAuthorizationTokenInput{
			RegistryIds: []string{},
		}
	}

	result, err := svc.GetAuthorizationTokenRequest(input).Send()
	if err != nil {
		if erroa, ok := err.(awserr.Error); ok {
			switch erroa.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, erroa.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, erroa.Error())
			default:
				fmt.Println(erroa.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.

			err = erroa
			fmt.Println(err.Error())
		}
	}

	if len(result.AuthorizationData) > 0 {
		token = aws.StringValue(result.AuthorizationData[0].AuthorizationToken)
	}

	return token, err
}

// SetRegistyID -
func (cr ECR) SetRegistyID(registryID string) {
	cr.registryID = registryID
}

// SetDefaultRegion -
func (cr ECR) SetDefaultRegion(defaultRegion string) {
	cr.defaultRegion = defaultRegion
}
