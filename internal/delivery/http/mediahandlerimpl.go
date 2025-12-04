package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaHandlerImpl struct {
	mediaApp           MediaApplication
	ErrRequiredMediaID string
	ErrInvalidMediaID  string
}

var _ MediaHandler = &MediaHandlerImpl{}

func ProvideMediaHandler(mediaApp MediaApplication) *MediaHandlerImpl {
	return &MediaHandlerImpl{
		mediaApp:           mediaApp,
		ErrRequiredMediaID: "media_id is required",
		ErrInvalidMediaID:  "invalid media_id",
	}
}

// ListMedia godoc
//
//	@Summary		List all media
//	@Description	Get all media with optional search
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//	@Param			queryParams	query	ListMediaRequestQueryParams	true	"Query parameters"
//	@Success		200			{object}	PaginationResponseDTO[MediaResponseDTO]
//	@Failure		500			{object}	Error
//	@Router			/media [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *MediaHandlerImpl) List(ctx *gin.Context) {
	paginate, err := createPaginationRequestDtoFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}
	search, _ := ctx.GetQuery("search")
	media, err := h.mediaApp.List(ctx.Request.Context(), ListMediaRequestDTO{
		QueryParams: &ListMediaRequestQueryParams{
			PaginationRequestDTO: *paginate,
			Search:               search,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, media)
}

// GetMedia godoc
//
//	@Summary		Get media by ID
//	@Description	Get media details by ID
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		GetMediaRequestPathParams	true	"Path parameters"
//	@Success		200			{object}	MediaResponseDTO
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/media/{id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *MediaHandlerImpl) Get(ctx *gin.Context) {
	mediaID, ok := pathToUUID(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidMediaID))
		return
	}

	media, err := h.mediaApp.Get(ctx.Request.Context(), GetMediaRequestDTO{
		PathParams: &GetMediaRequestPathParams{
			MediaID: mediaID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, media)
}

// CreateMedia godoc
//
//	@Summary		Create a new media
//	@Description	Create a new media
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//	@Param			media	body		CreateMediaRequestData	true	"Media request"
//	@Success		201		{object}	MediaResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/media [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *MediaHandlerImpl) Create(ctx *gin.Context) {
	var data CreateMediaRequestData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	media, err := h.mediaApp.Create(ctx.Request.Context(), CreateMediaRequestDTO{
		Data: &data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, media)
}

// DeleteMedia godoc
//
//	@Summary		Delete media
//	@Description	Delete media by ID
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		DeleteMediaRequestPathParams	true	"Path parameters"
//	@Success		204
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/media/{id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *MediaHandlerImpl) Delete(ctx *gin.Context) {
	mediaID, ok := pathToUUID(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidMediaID))
		return
	}

	err := h.mediaApp.Delete(ctx.Request.Context(), DeleteMediaRequestDTO{
		PathParams: &DeleteMediaRequestPathParams{
			MediaID: mediaID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
