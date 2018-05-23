# ECR Tool

[![Go Report Card](https://goreportcard.com/badge/github.com/lyon-pryde/ecr-tools)](https://goreportcard.com/report/github.com/lyon-pryde/ecr-tools)

Small container for generating docker token for ECR.

```sh
docker run pryde/ecr-tools get-token
```
Priority of config.

1) AWS Config file
2) Environment Variables
3) binary flags

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
