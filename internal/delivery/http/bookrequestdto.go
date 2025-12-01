package http

import (
	"github.com/google/uuid"
)

type ListBookRequestDTO struct{}

type CreateBookRequestDTO struct {
	Data *CreateBookRequestData
}

type CreateBookRequestData struct {
	Title         string      `json:"title"         binding:"required"`
	Description   string      `json:"description"   binding:"required"`
	Author        string      `json:"author"        binding:"required"`
	Price         int64       `json:"price"         binding:"required"`
	PagesCount    int         `json:"pagesCount"    binding:"required"`
	YearPublished int         `json:"yearPublished" binding:"required"`
	Publisher     string      `json:"publisher"     binding:"required"`
	Weight        float64     `json:"weight"        binding:"required"`
	StockQuantity int         `json:"stockQuantity" binding:"required"`
	CategoryIDs   []uuid.UUID `json:"categoryIds"   binding:"required"`
	Media         []uuid.UUID `json:"media"         binding:"required"`
}

type GetBookRequestDTO struct {
	ID uuid.UUID `json:"Id" binding:"required"`
}

type UpdateBookRequestDTO struct {
	ID   uuid.UUID `json:"Id" binding:"required"`
	Data *UpdateBookRequestData
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
	ID uuid.UUID `json:"Id" binding:"required"`
}
