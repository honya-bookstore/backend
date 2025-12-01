package http

type ListCategoryRequestDTO struct{}

type GetCategoryRequestDTO struct {
	ID string `json:"id" binding:"required"`
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
	ID   string `json:"id" binding:"required"`
	Data *UpdateCategoryRequestData
}

type UpdateCategoryRequestData struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteCategoryRequestDTO struct {
	ID string `json:"id" binding:"required"`
}
