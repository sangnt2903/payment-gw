package output

import "payment-gw/internal/core/domain"

type PaymentProvider interface {
	InitiatePayment(payment domain.Payment) (*domain.Payment, error)
	CheckPaymentStatus(paymentID string) (domain.PaymentStatus, error)
}
