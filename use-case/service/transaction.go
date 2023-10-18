package service

import (
	"github.com/facuellarg/stori/domain/entities"
	"github.com/facuellarg/stori/use-case/repository"
)

type TransactionService struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return TransactionService{repository}
}

func (t TransactionService) Store(transaction entities.Transaction) error {
	return t.TransactionRepository.Store(transaction)
}
