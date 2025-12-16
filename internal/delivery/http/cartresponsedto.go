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
	ID            uuid.UUID                      `json:"id"            binding:"required"`
	Title         string                         `json:"title"         binding:"required"`
	Description   string                         `json:"description"`
	Author        string                         `json:"author"        binding:"required"`
	Price         int64                          `json:"price"         binding:"required"`
	PagesCount    int                            `json:"pagesCount"    binding:"required"`
	YearPublished int                            `json:"yearPublished" binding:"required"`
	Publisher     string                         `json:"publisher"     binding:"required"`
	Weight        float64                        `json:"weight"`
	StockQuantity int                            `json:"stockQuantity" binding:"required"`
	PurchaseCount int                            `json:"purchaseCount" binding:"required"`
	Rating        float64                        `json:"rating"        binding:"required"`
	Medium        []CartItemBookMediaResponseDTO `json:"medium"        binding:"required"`
}

type CartItemBookMediaResponseDTO struct {
	ID      uuid.UUID `json:"id"      binding:"required"`
	IsCover bool      `json:"isCover" binding:"required"`
	Order   int       `json:"order"   binding:"required"`
	AltText string    `json:"altText"`
	URL     string    `json:"url"     binding:"required"`
}

func ToCartResponseDTO(cart *domain.Cart, bookMap map[uuid.UUID]*domain.Book, mediaMap map[uuid.UUID]*domain.Media) *CartResponseDTO {
	if cart == nil {
		return nil
	}

	items := make([]CartItemResponseDTO, 0, len(cart.Items))
	for _, item := range cart.Items {
		book := bookMap[item.BookID]
		if book == nil {
			continue
		}
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
				Medium:        ToCartItemBookMediaResponseDTOs(book.Medium, mediaMap),
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

func ToCartItemBookMediaResponseDTOs(media []domain.BookMedia, mediaMap map[uuid.UUID]*domain.Media) []CartItemBookMediaResponseDTO {
	dtos := make([]CartItemBookMediaResponseDTO, 0, len(media))
	for _, m := range media {
		mediaData := mediaMap[m.MediaID]
		if mediaData != nil {
			dtos = append(dtos, CartItemBookMediaResponseDTO{
				ID:      m.MediaID,
				IsCover: m.IsCover,
				Order:   m.Order,
				AltText: mediaData.AltText,
				URL:     mediaData.URL,
			})
		}
	}
	return dtos
}
