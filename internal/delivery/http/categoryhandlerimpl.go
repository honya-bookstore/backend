package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandlerImpl struct {
	categoryApp           CategoryApplication
	ErrRequiredCategoryID string
	ErrRequiredSlug       string
	ErrInvalidCategoryID  string
}

var _ CategoryHandler = &CategoryHandlerImpl{}

func ProvideCategoryHandler(categoryApp CategoryApplication) *CategoryHandlerImpl {
	return &CategoryHandlerImpl{
		categoryApp:           categoryApp,
		ErrRequiredCategoryID: "category_id is required",
		ErrRequiredSlug:       "slug is required",
		ErrInvalidCategoryID:  "invalid category_id",
	}
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
	paginate, err := createPaginationRequestDtoFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}
	search, _ := ctx.GetQuery("search")
	categories, err := h.categoryApp.List(ctx.Request.Context(), ListCategoryRequestDTO{
		QueryParams: &ListCategoryRequestQueryParams{
			PaginationRequestDTO: *paginate,
			Search:               search,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// GetCategoryBySlug godoc
//
//	@Summary		GetBySlug category by Slug
//	@Description	GetBySlug category details by Slug
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		GetCategoryBySlugRequestPathParams	true	"Path parameters"
//	@Success		200			{object}	CategoryResponseDTO
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/categories/{slug} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CategoryHandlerImpl) GetBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredSlug))
		return
	}

	category, err := h.categoryApp.GetBySlug(ctx.Request.Context(), GetCategoryBySlugRequestDTO{
		PathParams: &GetCategoryBySlugRequestPathParams{
			Slug: slug,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, category)
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
	var data CreateCategoryRequestData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	category, err := h.categoryApp.Create(ctx.Request.Context(), CreateCategoryRequestDTO{
		Data: &data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, category)
}

// UpdateCategory godoc
//
//	@Summary		Update category
//	@Description	Update category details
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		UpdateCategoryRequestPathParams	true	"Path parameters"
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
	categoryID, ok := pathToUUID(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCategoryID))
		return
	}

	var data UpdateCategoryRequestData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}
	updatedCategory, err := h.categoryApp.Update(ctx.Request.Context(), UpdateCategoryRequestDTO{
		PathParams: &UpdateCategoryRequestPathParams{
			CategoryID: categoryID,
		},
		Data: &data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory godoc
//
//	@Summary		Delete category
//	@Description	Delete category by ID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		DeleteCategoryRequestPathParams	true	"Path parameters"
//	@Success		204
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/categories/{id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CategoryHandlerImpl) Delete(ctx *gin.Context) {
	categoryID, ok := pathToUUID(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCategoryID))
		return
	}

	err := h.categoryApp.Delete(ctx.Request.Context(), DeleteCategoryRequestDTO{
		PathParams: &DeleteCategoryRequestPathParams{
			CategoryID: categoryID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
