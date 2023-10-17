package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DataLoader interface {
	Load() []Transaction
}

type Transformer interface {
	Transform([]string) (Transaction, error)
}

type CSVLoaderTransformer struct {
	FileName string
}

func (c CSVLoaderTransformer) Load() []Transaction {

	//Read csv file with headers
	file, err := os.OpenFile(c.FileName, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.TrimLeadingSpace = true

	//discard headers
	if _, err := reader.Read(); err != nil {
		panic(err)
	}

	response := []Transaction{}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		transaction, err := c.Transform(record)
		if err != nil {
			panic(err)
		}

		response = append(response, transaction)

	}
	return response

}

func (c CSVLoaderTransformer) Transform(record []string) (Transaction, error) {

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return Transaction{}, fmt.Errorf("invalid id %w", err)
	}

	dateSplited := strings.Split(record[1], "/")
	month, err := strconv.Atoi(dateSplited[0])
	if err != nil {
		return Transaction{}, fmt.Errorf("invalid month %w", err)
	}

	day, err := strconv.Atoi(dateSplited[1])
	if err != nil {
		return Transaction{}, fmt.Errorf("invalid day %w", err)
	}

	amount, err := strconv.ParseFloat(record[2], 32)
	if err != nil {
		return Transaction{}, fmt.Errorf("invalid amount %w", err)
	}

	return Transaction{
		ID:     id,
		Amount: float32(amount),
		Month:  month,
		Day:    day,
	}, nil

}
