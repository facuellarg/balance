package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func GetRegion() string {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1"
	}
	return region
}

func mySession() *session.Session {
	return session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(GetRegion()),
		}),
	)
}

func Dynamodb() *dynamodb.DynamoDB {
	return dynamodb.New(mySession())
}

type DynamoDBI interface {
	dynamodbiface.DynamoDBAPI
}
