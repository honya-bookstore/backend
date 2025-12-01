package http

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ListOrderRequestDTO struct{}

type GetOrderRequestDTO struct {
	ID uuid.UUID
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
	ID   uuid.UUID
	Data UpdateOrderData
}

type UpdateOrderData struct {
	Address string             `json:"address" binding:"required"`
	Status  domain.OrderStatus `json:"status"  binding:"required"`
	IsPaid  bool               `json:"isPaid"  binding:"required"`
}

type DeleteOrderRequestDTO struct {
	ID uuid.UUID
}
