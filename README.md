# ECR Tool

[![Go Report Card](https://goreportcard.com/badge/github.com/lyon-pryde/ecr-tools)](https://goreportcard.com/report/github.com/lyon-pryde/ecr-tools)

Small container for generating docker token for ECR.

```sh
docker run pryde/ecr-tools get-token
```

Priority of config.

'''sh
export AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
export AWS_ACCESS_KEY=$AWS_ACCESS_KEY
export AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY
export AWS_SECRET_KEY=AWS_SECRET_KEY
export AWS_DEFAULT_REGION
export AWS_ECR_REGISTRY_ID
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
