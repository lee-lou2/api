package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"log"
	"os"
)

func awsConfig() (aws.Config, error) {
	clientKey := os.Getenv("AWS_CLIENT_KEY")
	fmt.Println("클라이언트키")
	log.Println("클라이언트키")
	fmt.Println(clientKey)
	log.Println(clientKey)
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
	)
}
