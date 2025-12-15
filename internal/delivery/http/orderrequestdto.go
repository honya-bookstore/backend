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
	Status domain.OrderStatus `json:"status" enums:"pending,processing,shipping,delivered,cancelled"`
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
	Phone     string               `json:"phone"     binding:"required,e164"`
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
	IsPaid  bool               `json:"isPaid"`
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
	Amount            string `form:"vnp_Amount"            binding:"required"`
	BankTranNo        string `form:"vnp_BankTranNo"        binding:"required"`
	BankCode          string `form:"vnp_BankCode"          binding:"required"`
	CardType          string `form:"vnp_CardType"          binding:"required"`
	OrderInfo         string `form:"vnp_OrderInfo"         binding:"required"`
	PayDate           string `form:"vnp_PayDate"           binding:"required"`
	ResponseCode      string `form:"vnp_ResponseCode"      binding:"required"`
	SecureHash        string `form:"vnp_SecureHash"        binding:"required"`
	TmnCode           string `form:"vnp_TmnCode"           binding:"required"`
	TransactionNo     string `form:"vnp_TransactionNo"     binding:"required"`
	TransactionStatus string `form:"vnp_TransactionStatus" binding:"required"`
	TxnRef            string `form:"vnp_TxnRef"            binding:"required"`
}
