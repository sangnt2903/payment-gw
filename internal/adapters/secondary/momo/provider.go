package momo

import (
	"payment-gw/internal/core/domain"
	"payment-gw/internal/core/ports/output"
	"payment-gw/pkg/conf"

	uuid "github.com/satori/go.uuid"
)

type momoProvider struct {
	redirectURL string
}

func NewProvider() output.PaymentProvider {
	return &momoProvider{
		redirectURL: conf.GetString("momo", "redirect_url"),
	}
}

func (m *momoProvider) InitiatePayment(payment domain.Payment) (*domain.Payment, error) {
	payment.ProviderTransactionID = uuid.NewV4().String()
	return &payment, nil
}

func (m *momoProvider) CheckPaymentStatus(paymentID string) (domain.PaymentStatus, error) {
	return domain.PaymentStatusPending, nil
}
