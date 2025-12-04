package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandlerImpl struct {
	bookApp           BookApplication
	ErrRequiredBookID string
	ErrInvalidBookID  string
}

var _ BookHandler = &BookHandlerImpl{}

func ProvideBookHandler(
	bookApp BookApplication,
) *BookHandlerImpl {
	return &BookHandlerImpl{
		bookApp:           bookApp,
		ErrRequiredBookID: "book_id is required",
		ErrInvalidBookID:  "book_id is invalid",
	}
}

// ListBooks godoc
//
//	@Summary		List all books
//	@Description	Get all books with optional filters
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			queryParams	query	ListBookRequestQueryParams	true	"Query parameters"
//	@Success		200			{object}	PaginationResponseDTO[BookResponseDTO]
//	@Failure		500			{object}	Error
//	@Router			/books [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *BookHandlerImpl) List(ctx *gin.Context) {
	paginate, error := createPaginationRequestDtoFromQuery(ctx)
	if error != nil {
		return
	}
	search, _ := ctx.GetQuery("search")
	categoryIDs, _ := queryArrayToUUIDSlice(ctx, "category_ids")
	publisher, _ := ctx.GetQuery("publisher")
	yearString, _ := ctx.GetQuery("year")
	var year int
	if _, err := fmt.Sscanf(yearString, "%d", &year); err != nil {
		year = 0
	}
	var minPrice int64
	if minPriceQuery, ok := ctx.GetQuery("min_price"); ok {
		var price int64
		if _, err := fmt.Sscanf(minPriceQuery, "%d", &price); err == nil {
			minPrice = price
		}
	}
	var maxPrice int64
	if maxPriceQuery, ok := ctx.GetQuery("max_price"); ok {
		var price int64
		if _, err := fmt.Sscanf(maxPriceQuery, "%d", &price); err == nil {
			maxPrice = price
		}
	}
	sortRecent, _ := ctx.GetQuery("sort_recent")
	sortPrice, _ := ctx.GetQuery("sort_price")
	books, err := h.bookApp.List(ctx, ListBookRequestDTO{
		QueryParams: &ListBookRequestQueryParams{
			PaginationRequestDTO: *paginate,
			CategoryIDs:          categoryIDs,
			Search:               search,
			Publisher:            publisher,
			Year:                 year,
			MinPrice:             minPrice,
			MaxPrice:             maxPrice,
			SortRecent:           sortRecent,
			SortPrice:            sortPrice,
		},
	},
	)
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, books)
}

// GetBook godoc
//
//	@Summary		Get book by ID
//	@Description	Get book details by ID
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		GetBookRequestPathParams	true	"Path parameters"
//	@Success		200			{object}	BookResponseDTO
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/books/{id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *BookHandlerImpl) Get(ctx *gin.Context) {
	bookID, ok := pathToUUID(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidBookID))
		return
	}
	book, err := h.bookApp.Get(ctx, GetBookRequestDTO{
		PathParams: &GetBookRequestPathParams{
			BookID: bookID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, book)
}

// CreateBook godoc
//
//	@Summary		Create a new book
//	@Description	Create a new book
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			book	body		CreateBookData	true	"Book request"
//	@Success		201		{object}	BookResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/books [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *BookHandlerImpl) Create(ctx *gin.Context) {
	var data CreateBookData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	product, err := h.bookApp.Create(ctx.Request.Context(), CreateBookRequestDTO{
		Data: &data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, product)
}

// UpdateBook godoc
//
//	@Summary		Update book
//	@Description	Update book details
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		UpdateBookRequestPathParams	true	"Path parameters"
//	@Param			data	body		UpdateBookData	true	"Update book request"
//	@Success		200		{object}	BookResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		404		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/books/{id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *BookHandlerImpl) Update(ctx *gin.Context) {
	bookID, ok := pathToUUID(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidBookID))
		return
	}
	var data UpdateBookData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}
	updatedBook, err := h.bookApp.Update(ctx.Request.Context(), UpdateBookRequestDTO{
		PathParams: &UpdateBookRequestPathParams{
			BookID: bookID,
		},
		Data: &data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, updatedBook)
}

// DeleteBook godoc
//
//	@Summary		Delete book
//	@Description	Delete book by ID
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		DeleteBookRequestPathParams	true	"Path parameters"
//	@Success		204
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/books/{id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *BookHandlerImpl) Delete(ctx *gin.Context) {
	bookID, ok := pathToUUID(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidBookID))
		return
	}
	err := h.bookApp.Delete(ctx.Request.Context(), DeleteBookRequestDTO{
		PathParams: &DeleteBookRequestPathParams{
			BookID: bookID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
