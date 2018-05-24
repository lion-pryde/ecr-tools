# ECR Tool

[![Go Report Card](https://goreportcard.com/badge/github.com/lyon-pryde/ecr-tools)](https://goreportcard.com/report/github.com/lyon-pryde/ecr-tools)

## TL;DR

Small container for generating docker token for ECR.

```sh
docker run pryde/ecr-tools get-token
```

Priority of config.

- AWS Config files.

- - [~/.aws/config](~/.aws/config)

- - [~/.aws/credentials](~/.aws/credentials)

- Environment Variables
- Binary flags

## Required claims

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

## Examples

```sh
docker run pryde/ecr-tools get-token -i
```

```bash
bash-3.2$ go run main.go --help
  -aws-access-id string
        AWS Access ID overrides env variable AWS_ACCESS_KEY_ID
  -aws-default-region string
        AWS Default Region overrides env variable AWS_DEFAULT_REGION
  -aws-secret-access-key string
        AWS Access ID overrides env variable AWS_ACCESS_KEY_ID
  -i string
        AWS Access ID overrides env variable AWS_ACCESS_KEY_ID (shorthand)
  -k string
        AWS Access ID overrides env variable AWS_ACCESS_KEY_ID (shorthand)
  -r string
        AWS Default Region overrides env variable AWS_DEFAULT_REGION (shorthand)
```