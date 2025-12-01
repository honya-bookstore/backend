package http

import "github.com/google/uuid"

type ListArticleRequestDTO struct{}

type CreateArticleRequestDTO struct {
	Data *CreateArticleData
}

type CreateArticleData struct {
	Slug        string    `json:"slug"        binding:"required"`
	Title       string    `json:"title"       binding:"required"`
	ThumbnailID uuid.UUID `json:"thumbnailId" binding:"required"`
	Content     string    `json:"content"`
	Tags        []string  `json:"tags"`
}

type GetArticleRequestDTO struct {
	ID uuid.UUID `json:"Id" binding:"required"`
}

type UpdateArticleRequestDTO struct {
	ID   uuid.UUID          `json:"Id"   binding:"required"`
	Data *UpdateArticleData `json:"data"`
}

type UpdateArticleData struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	ThumbnailID uuid.UUID `json:"thumbnailId"`
	Content     string    `json:"content"`
}

type DeleteArticleRequestDTO struct {
	ID uuid.UUID `json:"Id" binding:"required"`
}
