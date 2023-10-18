package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/facuellarg/stori/external-service/controller"
)

type (
	TransactionRequest struct {
		FileName string `json:"fileName"`
		To       string `json:"to"`
	}

	TransactionServer struct {
		transactionController controller.TransactionController
	}
)

func NewTransactionServer(transactionController controller.TransactionController) TransactionServer {
	return TransactionServer{transactionController}
}

func (ts *TransactionServer) processTransaction(ctx context.Context, event *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	var transactionRequest TransactionRequest

	if err := json.NewDecoder(strings.NewReader(event.Body)).Decode(&transactionRequest); err != nil {
		fmt.Println(err)
		return &events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}
	if err := ts.transactionController.ProcessTransaction(
		transactionRequest.FileName,
		transactionRequest.To,
	); err != nil {
		fmt.Println(err)
		return &events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}
	return &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func (ts *TransactionServer) ProcessTransaction() {
	lambda.Start(ts.processTransaction)
}
