package application

import (
	"context"

	"backend/internal/domain"
)

type OrderPaymentService interface {
	GetPaymentURL(ctx context.Context, param GetPaymentURLParam) (string, error)
	VerifyIPN() error
}

type GetPaymentURLParam struct {
	ReturnURL string
	Order     *domain.Order
}
