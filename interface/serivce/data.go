package serivce

import "github.com/facuellarg/stori/domain/entities"

type DataLoader interface {
	Load(string) (entities.Transactions, error)
}

type Transformer interface {
	Transform([]string) (entities.Transaction, error)
}
