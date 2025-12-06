package application

import (
	"context"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type VNPayPaymentService interface {
	GetPaymentURL(
		ctx context.Context,
		param GetPaymentURLVNPayParam,
	) (string, error)

	VerifyIPN(
		ctx context.Context,
		param VerifyVNPayIPNParam,
		getOrder func(ctx context.Context, orderID uuid.UUID) (*domain.Order, error),
		onSuccess func(ctx context.Context, order *domain.Order) error,
		onFailure func(ctx context.Context, order *domain.Order) error,
	) (code string, message string)
}

type GetPaymentURLVNPayParam struct {
	ReturnURL string
	Order     *domain.Order
}

type VerifyVNPayIPNParam struct {
	OrderID uuid.UUID
	Data    *VerifyVNPayIPNData
}

type VerifyVNPayIPNData struct {
	Amount       string
	ResponseCode string
	SecureHash   string
	TmnCode      string
}
