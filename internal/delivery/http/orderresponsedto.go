package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type OrderResponseDTO struct {
	ID          uuid.UUID              `json:"id"                  binding:"required"`
	Address     string                 `json:"address"             binding:"required"`
	City        string                 `json:"city"                binding:"required"`
	Provider    domain.OrderProvider   `json:"provider"            binding:"required"`
	Status      domain.OrderStatus     `json:"status"              binding:"required"`
	IsPaid      bool                   `json:"isPaid"              binding:"required"`
	CreatedAt   time.Time              `json:"createdAt"           binding:"required"`
	UpdatedAt   time.Time              `json:"updatedAt"           binding:"required"`
	Items       []OrderItemResponseDTO `json:"items"               binding:"required"`
	TotalAmount int64                  `json:"totalAmount"         binding:"required"`
	UserID      uuid.UUID              `json:"userId"              binding:"required"`
	ReturnURL   string                 `json:"returnUrl,omitempty"`
}

type OrderItemResponseDTO struct {
	ID       uuid.UUID                 `json:"id"       binding:"required"`
	Book     *OrderItemBookResponseDTO `json:"book"     binding:"required"`
	Quantity int                       `json:"quantity" binding:"required,gt=0"`
	Price    int64                     `json:"price"    binding:"required,gt=0"`
}

type OrderItemBookResponseDTO struct {
	ID            uuid.UUID `json:"id"            binding:"required"`
	Title         string    `json:"title"         binding:"required"`
	Author        string    `json:"author"        binding:"required"`
	Price         int64     `json:"price"         binding:"required"`
	StockQuantity int       `json:"stockQuantity" binding:"required"`
	Rating        float64   `json:"rating"        binding:"required"`
}

func ToOrderResponseDTO(order *domain.Order, bookMap map[uuid.UUID]*domain.Book, returnURL string) *OrderResponseDTO {
	if order == nil {
		return nil
	}

	items := make([]OrderItemResponseDTO, 0, len(order.Items))
	for _, item := range order.Items {
		book := bookMap[item.BookID]
		itemDto := OrderItemResponseDTO{
			ID:       item.ID,
			Quantity: item.Quantity,
			Price:    item.Price,
		}

		if book != nil {
			itemDto.Book = &OrderItemBookResponseDTO{
				ID:            book.ID,
				Title:         book.Title,
				Author:        book.Author,
				Price:         book.Price,
				StockQuantity: book.StockQuantity,
				Rating:        book.Rating,
			}
		}

		items = append(items, itemDto)
	}

	return &OrderResponseDTO{
		ID:          order.ID,
		Address:     order.Address,
		City:        order.City,
		Provider:    order.Provider,
		Status:      order.Status,
		IsPaid:      order.IsPaid,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		Items:       items,
		TotalAmount: order.TotalAmount,
		UserID:      order.UserID,
		ReturnURL:   returnURL,
	}
}
