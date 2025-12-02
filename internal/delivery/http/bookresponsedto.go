package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type BookResponseDTO struct {
	ID            uuid.UUID                 `json:"id"            binding:"required"`
	Title         string                    `json:"title"         binding:"required"`
	Description   string                    `json:"description"`
	Author        string                    `json:"author"        binding:"required"`
	Price         int64                     `json:"price"         binding:"required"`
	PagesCount    int                       `json:"pagesCount"    binding:"required"`
	YearPublished int                       `json:"yearPublished" binding:"required"`
	Publisher     string                    `json:"publisher"     binding:"required"`
	Weight        float64                   `json:"weight"`
	StockQuantity int                       `json:"stockQuantity" binding:"required"`
	PurchaseCount int                       `json:"purchaseCount" binding:"required"`
	Rating        float64                   `json:"rating"        binding:"required"`
	Categories    []BookCategoryResponseDTO `json:"categories"    binding:"required"`
	Media         []BookMediaResponseDTO    `json:"media"         binding:"required"`
	CreatedAt     time.Time                 `json:"createdAt"     binding:"required"`
	UpdatedAt     time.Time                 `json:"updatedAt"     binding:"required"`
	DeletedAt     *time.Time                `json:"deletedAt"`
}

type BookMediaResponseDTO struct {
	ID       uuid.UUID `json:"id"       binding:"required"`
	IsCover  bool      `json:"isCover"  binding:"required"`
	FileName string    `json:"fileName" binding:"required"`
	FileURL  string    `json:"fileUrl"  binding:"required"`
	AltText  string    `json:"altText"`
}

type BookCategoryResponseDTO struct {
	ID          uuid.UUID `json:"id"          binding:"required"`
	Slug        string    `json:"slug"        binding:"required"`
	Name        string    `json:"name"        binding:"required"`
	Description string    `json:"description"`
}

func ToBookResponseDTO(book *domain.Book, categories []*domain.Category, mediaMap map[string]*domain.Media) *BookResponseDTO {
	if book == nil {
		return nil
	}

	categoryDtos := make([]BookCategoryResponseDTO, 0, len(categories))
	for _, cat := range categories {
		if cat != nil {
			categoryDtos = append(categoryDtos, BookCategoryResponseDTO{
				ID:          cat.ID,
				Slug:        cat.Slug,
				Name:        cat.Name,
				Description: cat.Description,
			})
		}
	}

	mediaDtos := make([]BookMediaResponseDTO, 0, len(book.Media))
	for _, bookMedia := range book.Media {
		if m, exists := mediaMap[bookMedia.MediaID.String()]; exists && m != nil {
			mediaDtos = append(mediaDtos, BookMediaResponseDTO{
				ID:       m.ID,
				IsCover:  bookMedia.IsCover,
				FileName: m.AltText,
				FileURL:  m.URL,
				AltText:  m.AltText,
			})
		}
	}

	var deletedAt *time.Time
	if !book.DeletedAt.IsZero() {
		deletedAt = &book.DeletedAt
	}

	return &BookResponseDTO{
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
		Categories:    categoryDtos,
		Media:         mediaDtos,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
		DeletedAt:     deletedAt,
	}
}
