package http

import "github.com/google/uuid"

type ListArticleRequestDTO struct {
	QueryParams *ListArticleRequestQueryParams
}

type ListArticleRequestQueryParams struct {
	PaginationRequestDTO
	Search string `json:"search"`
}

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
	PathParams *GetArticleRequestPathParams
}

type GetArticleRequestPathParams struct {
	ArticleID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}

type UpdateArticleRequestDTO struct {
	PathParams *UpdateArticleRequestPathParams
	Data       *UpdateArticleData
}

type UpdateArticleRequestPathParams struct {
	ArticleID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}

type UpdateArticleData struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	ThumbnailID uuid.UUID `json:"thumbnailId"`
	Content     string    `json:"content"`
}

type DeleteArticleRequestDTO struct {
	PathParams *DeleteArticleRequestPathParams
}

type DeleteArticleRequestPathParams struct {
	ArticleID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}
