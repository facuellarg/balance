package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/facuellarg/stori/domain/entities"
)

const (
	TableName = "transactions"
)

type TransactionDynamoRepository struct {
	awsSession dynamodbiface.DynamoDBAPI
}

func NewTransactionDynamoRepository(awsSession dynamodbiface.DynamoDBAPI) TransactionDynamoRepository {
	return TransactionDynamoRepository{awsSession}
}

func (t *TransactionDynamoRepository) Store(transaction entities.Transaction) error {

	av, err := dynamodbattribute.MarshalMap(transaction)
	if err != nil {
		return err
	}

	_, err = t.awsSession.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	})

	return err
}
