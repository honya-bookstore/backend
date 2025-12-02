package domain

import (
	"context"

	"github.com/google/uuid"
)

type CartRepository interface {
	Get(
		ctx context.Context,
		params CartRepositoryGetParam,
	) (*Cart, error)

	Save(
		ctx context.Context,
		params CartRepositorySaveParam,
	) error
}

type CartRepositoryGetParam struct {
	CartID uuid.UUID
	UserID uuid.UUID
}

type CartRepositorySaveParam struct {
	Cart Cart
}
