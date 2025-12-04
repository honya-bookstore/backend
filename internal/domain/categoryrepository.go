package domain

import (
	"context"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	List(
		ctx context.Context,
		params CategoryRepositoryListParam,
	) (*[]Category, error)

	Count(
		ctx context.Context,
	) (*int, error)

	Get(
		ctx context.Context,
		params CategoryRepositoryGetParam,
	) (*Category, error)

	Save(
		ctx context.Context,
		params CategoryRepositorySaveParam,
	) error
}

type CategoryRepositoryListParam struct {
	CategoryIDs []uuid.UUID
	Search      string
	Limit       int
	Offset      int
}

type CategoryRepositoryCountParam struct {
	CategoryIDs []uuid.UUID
}

type CategoryRepositoryGetParam struct {
	CategoryID   uuid.UUID
	CategorySlug string
}

type CategoryRepositorySaveParam struct {
	Category Category
}
