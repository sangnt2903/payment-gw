package http

import (
	"net/http"
	"payment-gw/internal/adapters/primary/http/response"
	"payment-gw/internal/core/domain"
	"payment-gw/internal/core/ports/input"

	uuid "github.com/satori/go.uuid"

	"payment-gw/internal/adapters/primary/http/request"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUseCase input.PaymentUseCase
}

func NewPaymentHandler(useCase input.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: useCase,
	}
}

// @Summary Create a new payment
// @Description Create a new payment transaction
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body request.CreatePaymentRequest true "Payment Request"
// @Success 201 {object} response.SuccessResponse{data=domain.Payment}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/payments [post]
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req request.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    "invalid_request",
			Message: err.Error(),
		})
		return
	}

	payment := domain.Payment{
		Amount:      req.Amount,
		Currency:    req.Currency,
		Description: req.Description,
		Provider:    req.Provider,
		Metadata:    req.Metadata,
		Status:      domain.PaymentStatusPending,
	}

	id := uuid.NewV4()
	payment.ID = id.String()

	// CreatePayment will use redirect URL from config based on provider
	result, err := h.paymentUseCase.CreatePayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    "internal_error",
			Message: "Failed to create payment",
		})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Data: result,
	})
}

// @Summary Get payment by ID
// @Description Get payment details by ID
// @Tags payments
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} response.SuccessResponse{data=domain.Payment}
// @Failure 404 {object} response.ErrorResponse
// @Router /api/v1/payments/{id} [get]
func (h *PaymentHandler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	// Validate UUID format
	if _, err := uuid.FromString(id); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    "invalid_id",
			Message: "Invalid payment ID format",
		})
		return
	}

	payment, err := h.paymentUseCase.GetPayment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    "not_found",
			Message: "Payment not found",
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: payment,
	})
}
