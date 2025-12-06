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
		param VerifyIPNVNPayParam,
		getOrder func(ctx context.Context, orderID uuid.UUID) (*domain.Order, error),
		onSuccess func(ctx context.Context, order *domain.Order) error,
		onFailure func(ctx context.Context, order *domain.Order) error,
	) (code, message string, err error)
}

type GetPaymentURLVNPayParam struct {
	ReturnURL string
	Order     *domain.Order
}

type VerifyIPNVNPayParam struct {
	Amount            string
	BankTranNo        string
	BankCode          string
	CardType          string
	OrderInfo         string
	PayDate           string
	ResponseCode      string
	SecureHash        string
	TmnCode           string
	TransactionNo     string
	TransactionStatus string
	TxnRef            string
}
