package http

import (
	"github.com/google/uuid"
)

type ListBookRequestDTO struct {
	QueryParams *ListBookRequestQueryParams
}

type ListBookRequestQueryParams struct {
	PaginationRequestDto
	Search     string    `json:"search"`
	CategoryID uuid.UUID `json:"category_id"`
	Publisher  string    `json:"publisher"`
	Year       int       `json:"year"`
	MinPrice   int64     `json:"min_price"`
	MaxPrice   int64     `json:"max_price"`
	SortRecent string    `json:"sort_recent" enums:"asc,desc"`
	SortPrice  string    `json:"sort_price"  enums:"asc,desc"`
}

type CreateBookRequestDTO struct {
	Data *CreateBookRequestData
}

type CreateBookRequestData struct {
	Title         string      `json:"title"         binding:"required"`
	Description   string      `json:"description"`
	Author        string      `json:"author"        binding:"required"`
	Price         int64       `json:"price"         binding:"required"`
	PagesCount    int         `json:"pagesCount"    binding:"required"`
	YearPublished int         `json:"yearPublished" binding:"required"`
	Publisher     string      `json:"publisher"     binding:"required"`
	Weight        float64     `json:"weight"`
	StockQuantity int         `json:"stockQuantity" binding:"required"`
	CategoryIDs   []uuid.UUID `json:"categoryIds"   binding:"required"`
	Media         []uuid.UUID `json:"media"         binding:"required"`
}

type GetBookRequestDTO struct {
	PathParams *GetBookRequestPathParams
}

type GetBookRequestPathParams struct {
	BookID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}

type UpdateBookRequestDTO struct {
	PathParams *UpdateBookRequestPathParams
	Data       *UpdateBookRequestData
}

type UpdateBookRequestPathParams struct {
	BookID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}

type UpdateBookRequestData struct {
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	Author        string      `json:"author"`
	Price         int64       `json:"price"`
	PagesCount    int         `json:"pagesCount"`
	YearPublished int         `json:"yearPublished"`
	Publisher     string      `json:"publisher"`
	Weight        float64     `json:"weight"`
	StockQuantity int         `json:"stockQuantity"`
	CategoryIDs   []uuid.UUID `json:"categoryIds"`
}

type DeleteBookRequestDTO struct {
	PathParams *DeleteBookRequestPathParams
}

type DeleteBookRequestPathParams struct {
	BookID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}
