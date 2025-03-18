package output

import "payment-gw/internal/core/domain"

type PaymentRepository interface {
	Save(payment domain.Payment) (*domain.Payment, error)
	FindByID(id string) (*domain.Payment, error)
	FindAll() ([]domain.Payment, error)
}
