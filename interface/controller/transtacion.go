package controller

import "github.com/facuellarg/stori/interface/serivce"

type (
	TransactionLambdaController struct {
		emailService       serivce.EmailService
		transactionService serivce.TransactionService
		loader             serivce.DataLoader
	}
)

func NewTransactionLambdaController(
	emailService serivce.EmailService,
	transactionService serivce.TransactionService,
	loader serivce.DataLoader,
) *TransactionLambdaController {
	return &TransactionLambdaController{
		emailService:       emailService,
		transactionService: transactionService,
		loader:             loader,
	}
}

func (t *TransactionLambdaController) ProcessTransaction(
	fileName string,
	to string,
) error {
	transactions, err := t.loader.Load(fileName)
	if err != nil {
		return err
	}

	err = t.transactionService.StoreTransactions(transactions)
	if err != nil {
		return err
	}
	balance := transactions.GetBalance()

	err = t.emailService.Send(to, balance)
	if err != nil {
		return err
	}

	return nil
}
