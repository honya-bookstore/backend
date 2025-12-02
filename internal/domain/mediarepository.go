package domain

import (
	"context"

	"github.com/google/uuid"
)

type MediaRepository interface {
	List(
		ctx context.Context,
		params MediaRepositoryListParam,
	) (*[]Media, error)

	Count(
		ctx context.Context,
	) (*int, error)

	Get(
		ctx context.Context,
		params MediaRepositoryGetParam,
	) (*Media, error)

	Save(
		ctx context.Context,
		params MediaRepositorySaveParam,
	) error
}

type MediaRepositoryListParam struct {
	MediaIDs []uuid.UUID
	Search   string
	Limit    int
	Offset   int
}

type MediaRepositoryCountParam struct {
	MediaIDs []uuid.UUID
	Search   string
}

type MediaRepositoryGetParam struct {
	MediaID uuid.UUID
}

type MediaRepositorySaveParam struct {
	Media Media
}
