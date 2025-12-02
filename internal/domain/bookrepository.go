package domain

import (
	"context"

	"github.com/google/uuid"
)

type BookRepository interface {
	List(
		ctx context.Context,
		params BookRepositoryListParam,
	) (*[]Category, error)

	Count(
		ctx context.Context,
	) (*int, error)

	Get(
		ctx context.Context,
		params BookRepositoryGetParam,
	) (*Category, error)

	Save(
		ctx context.Context,
		params BookCategorySaveParam,
	) error
}

type BookRepositoryListParam struct {
	BookIDs     []uuid.UUID
	Search      string
	CategoryIDs []uuid.UUID
	Publisher   string
	Year        int
	MinPrice    int64
	MaxPrice    int64
	SortRecent  string
	SortPrice   string
	Limit       int
	Offset      int
}

type BookRepositoryCountParam struct {
	BookIDs     []uuid.UUID
	Search      string
	CategoryIDs []uuid.UUID
	Publisher   string
	Year        int
	MinPrice    int64
	MaxPrice    int64
}

type BookRepositoryGetParam struct {
	BookID uuid.UUID
}

type BookCategorySaveParam struct {
	Book Book
}
