package domain

import (
	"context"

	"github.com/google/uuid"
)

type BookRepository interface {
	List(
		ctx context.Context,
		params BookRepositoryListParam,
	) (*[]Book, error)

	Count(
		ctx context.Context,
	) (*int, error)

	Get(
		ctx context.Context,
		params BookRepositoryGetParam,
	) (*Book, error)

	Save(
		ctx context.Context,
		params BookRepositorySaveParam,
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

type BookRepositorySaveParam struct {
	Book Book
}
