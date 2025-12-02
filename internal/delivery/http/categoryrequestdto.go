package http

type ListCategoryRequestDTO struct {
	QueryParams *ListCategoryRequestQueryParams
}

type ListCategoryRequestQueryParams struct {
	PaginationRequestDTO
}

type GetCategoryRequestDTO struct {
	PathParams *GetCategoryRequestPathParams
}

type GetCategoryRequestPathParams struct {
	CategoryID string `json:"id" binding:"required"`
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
	CategoryID string `json:"id" binding:"required"`
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
	CategoryID string `json:"id" binding:"required"`
}
