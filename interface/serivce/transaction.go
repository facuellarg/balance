package serivce

import "github.com/facuellarg/stori/domain/entities"

type TransactionService interface {
	StoreTransactions(transactions entities.Transactions) error
}
