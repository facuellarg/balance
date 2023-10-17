package repository

import (
	"github.com/facuellarg/stori/domain/entities"
)

type TransactionRepository interface {
	Store(entities.Transaction) error
}
