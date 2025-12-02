package http

import (
	"github.com/gin-gonic/gin"
)

type CategoryHandlerImpl struct{}

var _ CategoryHandler = &CategoryHandlerImpl{}

func ProvideCategoryHandler() *CategoryHandlerImpl {
	return &CategoryHandlerImpl{}
}

// ListCategories godoc
//
//	@Summary		List all categories
//	@Description	Get all categories
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			queryParams	query	ListCategoryRequestQueryParams	true	"Query parameters"
//	@Success		200			{object}	PaginationResponseDTO[CategoryResponseDTO]
//	@Failure		500			{object}	Error
//	@Router			/categories [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CategoryHandlerImpl) List(ctx *gin.Context) {
}

// GetCategory godoc
//
//	@Summary		Get category by ID
//	@Description	Get category details by ID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string	true	"Category ID"
//	@Success		200			{object}	CategoryResponseDTO
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/categories/{id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CategoryHandlerImpl) Get(ctx *gin.Context) {
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	Create a new category
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			category	body		CreateCategoryRequestData	true	"Category request"
//	@Success		201		{object}	CategoryResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/categories [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CategoryHandlerImpl) Create(ctx *gin.Context) {
}

// UpdateCategory godoc
//
//	@Summary		Update category
//	@Description	Update category details
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string							true	"Category ID"
//	@Param			data	body		UpdateCategoryRequestData	true	"Update category request"
//	@Success		200		{object}	CategoryResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		404		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/categories/{id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CategoryHandlerImpl) Update(ctx *gin.Context) {
}

// DeleteCategory godoc
//
//	@Summary		Delete category
//	@Description	Delete category by ID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Category ID"
//	@Success		204
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/categories/{id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CategoryHandlerImpl) Delete(ctx *gin.Context) {
}
