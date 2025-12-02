package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type OrderResponseDTO struct {
	ID          uuid.UUID              `json:"id"          binding:"required"`
	Address     string                 `json:"address"     binding:"required"`
	Provider    domain.OrderProvider   `json:"provider"    binding:"required"`
	Status      domain.OrderStatus     `json:"status"      binding:"required"`
	IsPaid      bool                   `json:"isPaid"      binding:"required"`
	CreatedAt   time.Time              `json:"createdAt"   binding:"required"`
	UpdatedAt   time.Time              `json:"updatedAt"   binding:"required"`
	Items       []OrderItemResponseDTO `json:"items"       binding:"omitempty,dive"`
	TotalAmount int64                  `json:"totalAmount" binding:"required"`
	UserID      uuid.UUID              `json:"userId"      binding:"required"`
}

type OrderItemResponseDTO struct {
	ID       uuid.UUID    `json:"id"       binding:"required"`
	Book     *domain.Book `json:"book"     binding:"required"`
	Quantity int          `json:"quantity" binding:"required,gt=0"`
	Price    int64        `json:"price"    binding:"required,gt=0"`
}

type OrderItemBookResponseDTO struct {
	ID            uuid.UUID `json:"id"            binding:"required"`
	Title         string    `json:"title"         binding:"required"`
	Description   string    `json:"description"`
	Author        string    `json:"author"`
	Price         int64     `json:"price"         binding:"required"`
	PagesCount    int       `json:"pagesCount"`
	YearPublished int       `json:"yearPublished"`
	Publisher     string    `json:"publisher"`
	Weight        float64   `json:"weight"`
	StockQuantity int       `json:"stockQuantity" binding:"required"`
	PurchaseCount int       `json:"purchaseCount" binding:"required"`
	Rating        float32   `json:"rating"        binding:"required"`
}
