package http

import (
	"time"

	"github.com/google/uuid"
)

type BookResponseDTO struct {
	ID            uuid.UUID                 `json:"id"            binding:"required"`
	Title         string                    `json:"title"         binding:"required"`
	Description   string                    `json:"description"`
	Author        string                    `json:"author"`
	Price         int64                     `json:"price"         binding:"required"`
	PagesCount    int                       `json:"pagesCount"`
	YearPublished int                       `json:"yearPublished"`
	Publisher     string                    `json:"publisher"`
	Weight        float64                   `json:"weight"`
	StockQuantity int                       `json:"stockQuantity" binding:"required"`
	PurchaseCount int                       `json:"purchaseCount" binding:"required"`
	Rating        float64                   `json:"rating"        binding:"required"`
	Categories    []BookCategoryResponseDTO `json:"categories"`
	Media         []BookMediaResponseDTO    `json:"media"`
	CreatedAt     time.Time                 `json:"createdAt"     binding:"required"`
	UpdatedAt     time.Time                 `json:"updatedAt"     binding:"required"`
	DeletedAt     *time.Time                `json:"deletedAt"`
}

type BookMediaResponseDTO struct {
	ID       uuid.UUID `json:"id"       binding:"required"`
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
