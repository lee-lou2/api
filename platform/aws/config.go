package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"os"
)

func awsConfig() (aws.Config, error) {
	clientKey := os.Getenv("AWS_CLIENT_KEY")
	secretKey := os.Getenv("AWS_CLIENT_SECRET")
	return config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				clientKey,
				secretKey,
				"",
			),
		),
		config.WithRegion("ap-northeast-2"),
	)
}
