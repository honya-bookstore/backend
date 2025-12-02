package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type CartResponseDTO struct {
	ID        uuid.UUID             `json:"id"        binding:"required"`
	Items     []CartItemResponseDTO `json:"items"     binding:"required"`
	UserID    uuid.UUID             `json:"userId"    binding:"required"`
	UpdatedAt time.Time             `json:"updatedAt" binding:"required"`
}

type CartItemResponseDTO struct {
	ID       uuid.UUID                `json:"id"       binding:"required"`
	Book     *CartItemBookResponseDTO `json:"book"     binding:"required"`
	Quantity int                      `json:"quantity" binding:"required"`
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
	Rating        float64   `json:"rating"        binding:"required"`
}

func ToCartResponseDTO(cart *domain.Cart, bookMap map[uuid.UUID]*domain.Book) *CartResponseDTO {
	if cart == nil {
		return nil
	}

	items := make([]CartItemResponseDTO, 0, len(cart.Items))
	for _, item := range cart.Items {
		book := bookMap[item.BookID]
		items = append(items, CartItemResponseDTO{
			ID: item.ID,
			Book: &CartItemBookResponseDTO{
				ID:            book.ID,
				Title:         book.Title,
				Description:   book.Description,
				Author:        book.Author,
				Price:         book.Price,
				PagesCount:    book.PagesCount,
				YearPublished: book.YearPublished,
				Publisher:     book.Publisher,
				Weight:        book.Weight,
				StockQuantity: book.StockQuantity,
				PurchaseCount: book.PurchaseCount,
				Rating:        book.Rating,
			},
			Quantity: item.Quantity,
		})
	}

	return &CartResponseDTO{
		ID:        cart.ID,
		Items:     items,
		UserID:    cart.UserID,
		UpdatedAt: cart.UpdatedAt,
	}
}
