package domain

import (
	"context"

	"github.com/google/uuid"
)

type OrderRepository interface {
	List(
		ctx context.Context,
		params OrderRepositoryListParam,
	) (*[]Order, error)

	Count(
		ctx context.Context,
	) (*int, error)

	Get(
		ctx context.Context,
		params OrderRepositoryGetParam,
	) (*Order, error)

	Save(
		ctx context.Context,
		params OrderRepositorySaveParam,
	) error
}

type OrderRepositoryListParam struct {
	OrderIDs []uuid.UUID
	Status   OrderStatus
	Limit    int
	Offset   int
}

type OrderRepositoryCountParam struct {
	OrderIDs    []uuid.UUID
	OrderStatus OrderStatus
}

type OrderRepositoryGetParam struct {
	OrderID uuid.UUID
}

type OrderRepositorySaveParam struct {
	Order Order
}
