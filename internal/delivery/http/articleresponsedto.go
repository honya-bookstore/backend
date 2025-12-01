package http

import (
	"time"

	"github.com/google/uuid"
)

type ArticleResponseDTO struct {
	ID        uuid.UUID                    `json:"id"        binding:"required"`
	Slug      string                       `json:"slug"      binding:"required"`
	Title     string                       `json:"title"     binding:"required"`
	Thumbnail *ArticleThumbnailResponseDTO `json:"thumbnail"`
	Content   *string                      `json:"content"`
	Tags      []string                     `json:"tags"`
	UserID    uuid.UUID                    `json:"userId"    binding:"required"`
	CreatedAt time.Time                    `json:"createdAt" binding:"required"`
	UpdatedAt time.Time                    `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time                   `json:"deletedAt"`
}

type ArticleThumbnailResponseDTO struct {
	ID       uuid.UUID `json:"id"       binding:"required"`
	FileName string    `json:"fileName" binding:"required"`
	FileURL  string    `json:"fileUrl"  binding:"required"`
	AltText  *string   `json:"altText"`
}
