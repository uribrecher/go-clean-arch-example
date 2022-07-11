package payment

import "clean/internal/entities"

type Service interface {
	CreateTransaction(amount string, from, to entities.Payment) string
	ConfirmTransaction(txToken string) error
}
