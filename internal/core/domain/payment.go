package domain

import (
	"time"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSuccess   PaymentStatus = "success"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

type Payment struct {
	ID                   string                 `json:"id"`
	Amount               float64                `json:"amount"`
	Currency             string                 `json:"currency"`
	Status               PaymentStatus          `json:"status"`
	Description          string                 `json:"description"`
	Provider            string                 `json:"provider"`
	ProviderTransactionID string                 `json:"provider_transaction_id,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt            time.Time              `json:"created_at"`
	UpdatedAt            time.Time              `json:"updated_at"`
}