package application

import (
	"context"

	"backend/internal/domain"
)

type VNPayPaymentService interface {
	GetPaymentURL(ctx context.Context, param GetPaymentURLVNPayParam) (string, error)
	VerifyIPN(ctx context.Context, param VerifyVNPayIPNParam) (code string, message string)
}

type GetPaymentURLVNPayParam struct {
	ReturnURL string
	Order     *domain.Order
}
type VerifyVNPayIPNParam struct {
	Order *domain.Order
	Data  *VerifyVNPayIPNData
}

type VerifyVNPayIPNData struct {
	Amount       string
	ResponseCode string
	SecureHash   string
	TmnCode      string
}
