package http

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ListOrderRequestDTO struct {
	QueryParams *ListOrderRequestQueryParams
}

type ListOrderRequestQueryParams struct {
	PaginationRequestDTO
	Status domain.OrderStatus `json:"status" enums:"pending,processing,shipped,delivered,cancelled"`
}

type GetOrderRequestDTO struct {
	PathParams *GetOrderRequestPathParams
}

type GetOrderRequestPathParams struct {
	OrderID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}

type CreateOrderRequestDTO struct {
	Data *CreateOrderData
}

type CreateOrderData struct {
	Email     string               `json:"email"     binding:"required,email"`
	FirstName string               `json:"firstName" binding:"required"`
	LastName  string               `json:"lastName"  binding:"required"`
	Address   string               `json:"address"   binding:"required"`
	City      string               `json:"city"      binding:"required"`
	Provider  domain.OrderProvider `json:"provider"  binding:"required"`
	UserID    uuid.UUID            `json:"userId"    binding:"required"`
	ReturnURL string               `json:"returnUrl"`
}

type UpdateOrderRequestDTO struct {
	PathParams *UpdateOrderRequestPathParams
	Data       UpdateOrderData
}

type UpdateOrderRequestPathParams struct {
	OrderID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}

type UpdateOrderData struct {
	Address string             `json:"address" binding:"required"`
	Status  domain.OrderStatus `json:"status"  binding:"required"`
	IsPaid  bool               `json:"isPaid"  binding:"required"`
}

type DeleteOrderRequestDTO struct {
	PathParams *DeleteOrderRequestPathParams
}

type DeleteOrderRequestPathParams struct {
	OrderID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}

type VerifyVNPayIPNRequestDTO struct {
	QueryParams *VerifyVNPayIPNQueryParams
}

type VerifyVNPayIPNQueryParams struct {
	Amount            string `json:"vnp_Amount"`
	BackTranNo        string `json:"vnp_BankTranNo"`
	BankCode          string `json:"vnp_BankCode"`
	CartType          string `json:"vnp_CardType"`
	OrderInfo         string `json:"vnp_OrderInfo"`
	PayDate           string `json:"vnp_PayDate"`
	ResponseCode      string `json:"vnp_ResponseCode"`
	SecureHash        string `json:"vnp_SecureHash"`
	TmnCode           string `json:"vnp_TmnCode"`
	TransactionNo     string `json:"vnp_TransactionNo"`
	TransactionStatus string `json:"vnp_TransactionStatus"`
	TxnRef            string `json:"vnp_TxnRef"`
}
