package services

import (
	"payment-gw/internal/core/domain"
	"payment-gw/internal/core/ports/input"
	"payment-gw/internal/core/ports/output"
	"time"
)

type paymentService struct {
	paymentRepo     output.PaymentRepository
	paymentProvider output.PaymentProvider
}

func NewPaymentService(repo output.PaymentRepository, provider output.PaymentProvider) input.PaymentUseCase {
	return &paymentService{
		paymentRepo:     repo,
		paymentProvider: provider,
	}
}

func (s *paymentService) CreatePayment(payment domain.Payment) (*domain.Payment, error) {
	// Initialize payment with pending status
	payment.Status = domain.PaymentStatusPending
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	// Initiate payment with Momo
	processedPayment, err := s.paymentProvider.InitiatePayment(payment)
	if err != nil {
		return nil, err
	}

	// Save payment to database
	return s.paymentRepo.Save(*processedPayment)
}

func (s *paymentService) GetPayment(id string) (*domain.Payment, error) {
	return s.paymentRepo.FindByID(id)
}

func (s *paymentService) ListPayments() ([]domain.Payment, error) {
	return s.paymentRepo.FindAll()
}
