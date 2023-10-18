package service

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/facuellarg/stori/domain/entities"
)

type CSVLoaderTransformer struct {
	s3     s3iface.S3API
	bucket string
}

func NewCSVLoaderTransformer(s3 s3iface.S3API, bucket string) CSVLoaderTransformer {
	return CSVLoaderTransformer{s3, bucket}
}

func (c CSVLoaderTransformer) Load(fileName string) (entities.Transactions, error) {

	input := &s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(fileName),
	}

	result, err := c.s3.GetObject(input)
	if err != nil {
		err = fmt.Errorf("error getting object %s from bucket %s: %w", fileName, c.bucket, err)
		return nil, err
	}

	defer result.Body.Close()

	reader := csv.NewReader(result.Body)
	reader.FieldsPerRecord = 3
	reader.TrimLeadingSpace = true

	//discard headers
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	response := []entities.Transaction{}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		transaction, err := c.Transform(record)
		if err != nil {
			return nil, err
		}

		response = append(response, transaction)

	}
	return response, nil

}

func (c CSVLoaderTransformer) Transform(record []string) (entities.Transaction, error) {

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return entities.Transaction{}, fmt.Errorf("invalid id %w", err)
	}

	dateSplited := strings.Split(record[1], "/")
	month, err := strconv.Atoi(dateSplited[0])
	if err != nil {
		return entities.Transaction{}, fmt.Errorf("invalid month %w", err)
	}

	day, err := strconv.Atoi(dateSplited[1])
	if err != nil {
		return entities.Transaction{}, fmt.Errorf("invalid day %w", err)
	}

	amount, err := strconv.ParseFloat(record[2], 32)
	if err != nil {
		return entities.Transaction{}, fmt.Errorf("invalid amount %w", err)
	}

	return entities.Transaction{
		ID:     id,
		Amount: float32(amount),
		Month:  month,
		Day:    day,
	}, nil

}
