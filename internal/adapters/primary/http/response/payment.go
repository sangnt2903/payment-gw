package response

type PaymentResponse struct {
    ID                   string                 `json:"id"`
    Amount               float64                `json:"amount"`
    Currency             string                 `json:"currency"`
    Status               string                 `json:"status"`
    Description          string                 `json:"description"`
    Provider             string                 `json:"provider"`
    ProviderTransactionID string                 `json:"provider_transaction_id,omitempty"`
    PaymentURL           string                 `json:"payment_url,omitempty"`
    QRCode              string                 `json:"qr_code,omitempty"`
    ExpiryTime          int64                  `json:"expiry_time,omitempty"`
    Metadata            map[string]interface{} `json:"metadata,omitempty"`
    CreatedAt           string                 `json:"created_at"`
    UpdatedAt           string                 `json:"updated_at"`
}