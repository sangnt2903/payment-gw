package input

import "payment-gw/internal/core/domain"

type PaymentUseCase interface {
    CreatePayment(payment domain.Payment) (*domain.Payment, error)
    GetPayment(id string) (*domain.Payment, error)
    ListPayments() ([]domain.Payment, error)
}