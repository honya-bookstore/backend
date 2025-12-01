package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

// CartResponseDTO represents the response structure for a cart
type CartResponseDTO struct {
	ID        uuid.UUID             `json:"id"        binding:"required"`
	Items     []CartItemResponseDTO `json:"items"     binding:"required"`
	UserID    uuid.UUID             `json:"userId"    binding:"required"`
	UpdatedAt time.Time             `json:"updatedAt" binding:"required"`
}

// CartItemResponseDTO represents the response structure for a cart item
type CartItemResponseDTO struct {
	ID       uuid.UUID    `json:"id"       binding:"required"`
	Book     *domain.Book `json:"book"     binding:"required"`
	Quantity int          `json:"quantity" binding:"required"`
}

type CartItemBookResponseDTO struct {
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
