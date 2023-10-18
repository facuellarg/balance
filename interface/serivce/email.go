package serivce

import "github.com/facuellarg/stori/domain/entities"

type EmailService interface {
	Send(string, entities.Balance) error
}
