package http

import "github.com/google/uuid"

type ListCategoryRequestDTO struct {
	QueryParams *ListCategoryRequestQueryParams
}

type ListCategoryRequestQueryParams struct {
	PaginationRequestDTO
	Search string `json:"search"`
}

type GetCategoryBySlugRequestDTO struct {
	PathParams *GetCategoryBySlugRequestPathParams
}

type GetCategoryBySlugRequestPathParams struct {
	Slug string `json:"slug" binding:"required"`
}

type CreateCategoryRequestDTO struct {
	Data *CreateCategoryRequestData
}

type CreateCategoryRequestData struct {
	Slug        string `json:"slug"        binding:"required"`
	Name        string `json:"name"        binding:"required"`
	Description string `json:"description"`
}

type UpdateCategoryRequestDTO struct {
	PathParams *UpdateCategoryRequestPathParams
	Data       *UpdateCategoryRequestData
}

type UpdateCategoryRequestPathParams struct {
	CategoryID uuid.UUID `json:"id" binding:"required"`
}

type UpdateCategoryRequestData struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteCategoryRequestDTO struct {
	PathParams *DeleteCategoryRequestPathParams
}

type DeleteCategoryRequestPathParams struct {
	CategoryID uuid.UUID `json:"id" binding:"required"`
}
