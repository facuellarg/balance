package service

import (
	"fmt"
	"sync"

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

func (t TransactionService) StoreTransactions(transactions entities.Transactions) error {
	total := len(transactions)
	var wg sync.WaitGroup
	wg.Add(total)

	errChan := make(chan error, total)

	for _, transaction := range transactions {
		go func(transaction entities.Transaction, errChan chan error, wg *sync.WaitGroup) {
			defer wg.Done()
			if err := t.Store(transaction); err != nil {
				errChan <- err
			}
		}(transaction, errChan, &wg)
	}
	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return fmt.Errorf("error saving transaction: %w", err)
		}
	}

	return nil
}
