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
	Address  string                `json:"address"  binding:"required"`
	Provider domain.OrderProvider  `json:"provider" binding:"required"`
	Items    []CreateOrderItemData `json:"items"    binding:"required,dive"`
	UserID   uuid.UUID             `json:"userId"   binding:"required"`
}

type CreateOrderItemData struct {
	BookID   uuid.UUID `json:"bookId"   binding:"required"`
	Quantity int       `json:"quantity" binding:"required"`
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
