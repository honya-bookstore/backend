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
	ID       uuid.UUID `json:"id"       binding:"required"  validate:"required"`
	BookID   uuid.UUID `json:"bookID"   validate:"required"`
	Quantity int       `json:"quantity" binding:"required"  validate:"required,gt=0"`
	Price    int64     `json:"price"    binding:"required"  validate:"required,gt=0"`
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

func NewOrder(
	userID uuid.UUID,
	address string,
	provider OrderProvider,
	items []OrderItem,
) (*Order, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	totalAmount := int64(0)
	for _, item := range items {
		totalAmount += item.Price * int64(item.Quantity)
	}
	now := time.Now()
	return &Order{
		ID:          id,
		UserID:      userID,
		Address:     address,
		Provider:    provider,
		Status:      OrderStatusPending,
		IsPaid:      false,
		CreatedAt:   now,
		UpdatedAt:   now,
		Items:       items,
		TotalAmount: totalAmount,
	}, nil
}

func NewOrderItem(
	bookID uuid.UUID,
	quantity int,
	price int64,
) (*OrderItem, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &OrderItem{
		ID:       id,
		BookID:   bookID,
		Quantity: quantity,
		Price:    price,
	}, nil
}

func (o *Order) Update(
	address string,
	status OrderStatus,
	isPaid bool,
) {
	updated := false
	if o.Address != address {
		o.Address = address
		updated = true
	}
	if o.Status != status {
		o.Status = status
		updated = true
	}
	if o.IsPaid != isPaid {
		o.IsPaid = isPaid
		updated = true
	}
	if updated {
		o.UpdatedAt = time.Now()
	}
}
