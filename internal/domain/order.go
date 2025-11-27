package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID     `json:"id"          binding:"required"        validate:"required"`
	Address     string        `json:"address"     binding:"required"        validate:"required"`
	Provider    OrderProvider `json:"provider"    binding:"required"        validate:"required,oneof=COD VNPAY MOMO ZALOPAY"`
	Status      OrderStatus   `json:"status"      binding:"required"        validate:"required,oneof=Pending Processing Shipped Delivered Cancelled"`
	IsPaid      bool          `json:"isPaid"      binding:"required"        validate:"required"`
	CreatedAt   time.Time     `json:"createdAt"   binding:"required"        validate:"required"`
	UpdatedAt   time.Time     `json:"updatedAt"   binding:"required"        validate:"required,gtefield=CreatedAt"`
	Items       []OrderItem   `json:"items"       validate:"omitempty,dive"`
	TotalAmount int64         `json:"totalAmount" binding:"required"        validate:"required"`
	UserID      uuid.UUID     `json:"userId"      binding:"required"        validate:"required"`
}

type OrderItem struct {
	ID       uuid.UUID `json:"id"       binding:"required" validate:"required"`
	Book     *Book     `json:"book"`
	Quantity int       `json:"quantity" binding:"required" validate:"required,gt=0"`
	Price    int64     `json:"price"    binding:"required" validate:"required,gt=0"`
}

type OrderProvider string

const (
	PaymentProviderCOD     OrderProvider = "COD"
	PaymentProviderVNPAY   OrderProvider = "VNPAY"
	PaymentProviderMOMO    OrderProvider = "MOMO"
	PaymentProviderZALOPAY OrderProvider = "ZALOPAY"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "Pending"
	OrderStatusProcessing OrderStatus = "Processing"
	OrderStatusShipped    OrderStatus = "Shipped"
	OrderStatusDelivered  OrderStatus = "Delivered"
	OrderStatusCancelled  OrderStatus = "Cancelled"
)
