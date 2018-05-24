# ECR Tool

[![Go Report Card](https://goreportcard.com/badge/github.com/lyon-pryde/ecr-tools)](https://goreportcard.com/report/github.com/lyon-pryde/ecr-tools)

Small container for generating docker token for ECR.

```sh
docker run pryde/ecr-tools get-token
```
Priority of config.

'''sh
export AWSAccessKeyIDEnvVar = "AWS_ACCESS_KEY_ID"
export AWSAccessKeyEnvVar   = "AWS_ACCESS_KEY"

export AWSSecreteAccessKeyEnvVar = "AWS_SECRET_ACCESS_KEY"
export AWSSecreteKeyEnvVar       = "AWS_SECRET_KEY"
'''

> Required claims

```json
[
    "ecr:GetAuthorizationToken",
    "ecr:BatchCheckLayerAvailability",
    "ecr:GetDownloadUrlForLayer",
    "ecr:GetRepositoryPolicy",
    "ecr:DescribeRepositories",
    "ecr:ListImages",
    "ecr:DescribeImages",
    "ecr:BatchGetImage",
    "ecr:InitiateLayerUpload",
    "ecr:UploadLayerPart",
    "ecr:CompleteLayerUpload",
    "ecr:PutImage"
]
```
