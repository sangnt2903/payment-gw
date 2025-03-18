package momo

import (
	"payment-gw/internal/core/domain"
	"payment-gw/internal/core/ports/output"
	"payment-gw/pkg/conf"

	momosdk "github.com/momo-wallet/payment-sdk"
	uuid "github.com/satori/go.uuid"
)

type momoProvider struct {
	client      *momosdk.Client
	redirectURL string
}

func NewProvider() output.PaymentProvider {
	client := momosdk.NewClient(&momosdk.Config{
		PartnerCode: conf.GetString("momo", "partner_code"),
		AccessKey:   conf.GetString("momo", "access_key"),
		SecretKey:   conf.GetString("momo", "secret_key"),
		Environment: conf.GetString("momo", "environment"),
	})

	return &momoProvider{
		client:      client,
		redirectURL: conf.GetString("momo", "redirect_url"),
	}
}

func (m *momoProvider) InitiatePayment(payment domain.Payment) (*domain.Payment, error) {
	request := &momosdk.PaymentRequest{
		OrderID:     payment.ID,
		Amount:      int64(payment.Amount),
		OrderInfo:   payment.Description,
		RedirectURL: m.redirectURL,
		RequestID:   uuid.NewV4().String(),
		RequestType: "captureWallet",
		ExtraData:   "",
		Language:    "vi",
	}

	response, err := m.client.CreatePayment(request)
	if err != nil {
		return nil, err
	}

	payment.ProviderTransactionID = response.RequestID
	payment.PaymentURL = response.PayURL
	payment.QRCode = response.QRCode
	payment.ExpiryTime = response.ExpiryTime

	return &payment, nil
}

func (m *momoProvider) CheckPaymentStatus(paymentID string) (domain.PaymentStatus, error) {
	response, err := m.client.QueryPayment(&momosdk.QueryRequest{
		OrderID: paymentID,
	})
	if err != nil {
		return domain.PaymentStatusFailed, err
	}

	switch response.Status {
	case 0:
		return domain.PaymentStatusPending, nil
	case 1:
		return domain.PaymentStatusSuccess, nil
	default:
		return domain.PaymentStatusFailed, nil
	}
}
