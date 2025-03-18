package request

type CreatePaymentRequest struct {
    Amount      float64                `json:"amount" binding:"required"`
    Currency    string                 `json:"currency" binding:"required,len=3"`
    Description string                 `json:"description"`
    Provider    string                 `json:"provider" binding:"required,oneof=momo zalopay"`
    Metadata    map[string]interface{} `json:"metadata"`
}