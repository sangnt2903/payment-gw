package postgres

import (
    "context"
    "payment-gw/internal/core/domain"
    "payment-gw/internal/core/ports/output"

    "github.com/jackc/pgx/v5/pgxpool"
)

type paymentRepository struct {
    pool *pgxpool.Pool
}

func NewPaymentRepository(pool *pgxpool.Pool) output.PaymentRepository {
    return &paymentRepository{
        pool: pool,
    }
}

func (r *paymentRepository) Save(payment domain.Payment) (*domain.Payment, error) {
    query := `
        INSERT INTO payments (id, amount, currency, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`
    
    err := r.pool.QueryRow(
        context.Background(),
        query,
        payment.ID,
        payment.Amount,
        payment.Currency,
        payment.Status,
        payment.CreatedAt,
        payment.UpdatedAt,
    ).Scan(&payment.ID)

    if err != nil {
        return nil, err
    }
    
    return &payment, nil
}

func (r *paymentRepository) FindByID(id string) (*domain.Payment, error) {
    payment := &domain.Payment{}
    query := `
        SELECT id, amount, currency, status, created_at, updated_at
        FROM payments
        WHERE id = $1`
    
    err := r.pool.QueryRow(
        context.Background(),
        query,
        id,
    ).Scan(
        &payment.ID,
        &payment.Amount,
        &payment.Currency,
        &payment.Status,
        &payment.CreatedAt,
        &payment.UpdatedAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    return payment, nil
}

func (r *paymentRepository) FindAll() ([]domain.Payment, error) {
    query := `
        SELECT id, amount, currency, status, created_at, updated_at
        FROM payments`
    
    rows, err := r.pool.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var payments []domain.Payment
    for rows.Next() {
        var payment domain.Payment
        err := rows.Scan(
            &payment.ID,
            &payment.Amount,
            &payment.Currency,
            &payment.Status,
            &payment.CreatedAt,
            &payment.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        payments = append(payments, payment)
    }
    
    return payments, nil
}