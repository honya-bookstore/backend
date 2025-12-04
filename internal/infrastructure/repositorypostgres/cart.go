package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Cart struct {
	queries *sqlc.Queries
	conn    *pgxpool.Pool
}

var _ domain.CartRepository = (*Cart)(nil)

func ProvideCart(
	queries *sqlc.Queries,
	conn *pgxpool.Pool,
) *Cart {
	return &Cart{
		queries: queries,
		conn:    conn,
	}
}

func (r *Cart) Get(
	ctx context.Context,
	params domain.CartRepositoryGetParam,
) (*domain.Cart, error) {
	cartEntity, err := r.queries.GetCart(ctx, sqlc.GetCartParams{
		ID:     params.CartID,
		UserID: params.UserID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	cart := toCartDomain(&cartEntity)
	cartItems, err := r.queries.ListCartItems(ctx, sqlc.ListCartItemsParams{
		CartID: cart.ID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	cart.Items = make([]domain.CartItem, 0, len(cartItems))
	for _, cartItemEntity := range cartItems {
		cart.Items = append(cart.Items, domain.CartItem{
			ID:       cartItemEntity.ID,
			BookID:   cartItemEntity.BookID,
			Quantity: int(cartItemEntity.Quantity),
		})
	}
	return cart, nil
}

func (r *Cart) Save(
	ctx context.Context,
	params domain.CartRepositorySaveParam,
) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return toDomainError(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	qtx := r.queries.WithTx(tx)
	err = qtx.UpsertCart(ctx, sqlc.UpsertCartParams{
		ID:     params.Cart.ID,
		UserID: params.Cart.UserID,
		UpdatedAt: pgtype.Timestamptz{
			Time:  params.Cart.UpdatedAt,
			Valid: true,
		},
	})
	if err != nil {
		return toDomainError(err)
	}
	err = r.mergeCartItems(ctx, qtx, &params.Cart)
	if err != nil {
		return toDomainError(err)
	}
	if err = tx.Commit(ctx); err != nil {
		return toDomainError(err)
	}
	return nil
}

func toCartDomain(cartEntity *sqlc.Cart) *domain.Cart {
	return &domain.Cart{
		ID:     cartEntity.ID,
		UserID: cartEntity.UserID,
	}
}

func (r *Cart) mergeCartItems(
	ctx context.Context,
	qtx *sqlc.Queries,
	cart *domain.Cart,
) error {
	if err := qtx.CreateTempTableCartItems(ctx); err != nil {
		return err
	}
	param := make([]sqlc.InsertTempTableCartItemsParams, 0, len(cart.Items))
	for _, item := range cart.Items {
		param = append(param, sqlc.InsertTempTableCartItemsParams{
			ID:       item.ID,
			Quantity: int32(item.Quantity),
			CartID:   cart.ID,
			BookID:   item.BookID,
		})
	}
	if _, err := qtx.InsertTempTableCartItems(ctx, param); err != nil {
		return err
	}
	if err := qtx.MergeCartItemsFromTemp(ctx); err != nil {
		return err
	}
	return nil
}
