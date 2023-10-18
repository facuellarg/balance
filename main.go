package main

import (
	"fmt"
	"os"

	"github.com/facuellarg/stori/aws"
	"github.com/facuellarg/stori/external-service/server"
	"github.com/facuellarg/stori/interface/controller"
	"github.com/facuellarg/stori/interface/repository"
	"github.com/facuellarg/stori/use-case/service"
)

type Transaction struct {
	ID     int
	Amount float32
	Month  int
	Day    int
}

type TransactionMonthInfo struct {
	Total         int
	CreditAverage float32
	DebitAverage  float32
}

type Balance struct {
	TotalBalance         float32
	TransactionsPerMonth map[string]TransactionMonthInfo
	AverageDebitAmount   float32
	AverageCreditAmount  float32
}

var (
	pass   = os.Getenv("STORI_PASSWORD")
	email  = os.Getenv("STORI_EMAIL")
	bucket = os.Getenv("STORI_BUCKET")
)

func main() {

	fmt.Println("bucket: ", bucket)

	loader := service.NewCSVLoaderTransformer(aws.S3(), bucket)
	dynamoRepository := repository.NewTransactionDynamoRepository(aws.Dynamodb())
	transactionService := service.NewTransactionService(&dynamoRepository)

	emailService := service.NewEmailService(email, pass)
	transactionController := controller.NewTransactionLambdaController(emailService, transactionService, loader)
	// if err := transactionController.ProcessTransaction("transactions.csv", bucket, to); err != nil {
	// 	panic(err)
	// }
	transactionServer := server.NewTransactionServer(transactionController)
	transactionServer.ProcessTransaction()

}
