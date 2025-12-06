package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Order struct {
	queries *sqlc.Queries
	conn    *pgxpool.Pool
}

var _ domain.OrderRepository = (*Order)(nil)

func ProvideOrder(
	queries *sqlc.Queries,
	conn *pgxpool.Pool,
) *Order {
	return &Order{
		queries: queries,
		conn:    conn,
	}
}

func (r *Order) List(
	ctx context.Context,
	params domain.OrderRepositoryListParam,
) (*[]domain.Order, error) {
	var statusID uuid.UUID
	if params.Status != "" {
		statusEntity, err := r.queries.GetOrderStatus(ctx, sqlc.GetOrderStatusParams{
			ID:   uuid.Nil,
			Name: string(params.Status),
		})
		if err != nil {
			return nil, toDomainError(err)
		}
		statusID = statusEntity.ID
	}

	orderEntities, err := r.queries.ListOrders(ctx, sqlc.ListOrdersParams{
		IDs:      params.OrderIDs,
		StatusID: statusID,
		Offset:   int32(params.Offset),
		Limit:    int32(params.Limit),
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	if len(orderEntities) == 0 {
		emptyOrders := []domain.Order{}
		return &emptyOrders, nil
	}

	orderIDs := make([]uuid.UUID, 0, len(orderEntities))
	for _, orderEntity := range orderEntities {
		orderIDs = append(orderIDs, orderEntity.ID)
	}

	orderItemsEntities, err := r.queries.ListOrderItems(ctx, sqlc.ListOrderItemsParams{
		OrderIDs: orderIDs,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	statusEntities, err := r.queries.ListOrderStatuses(ctx, sqlc.ListOrderStatusesParams{})
	if err != nil {
		return nil, toDomainError(err)
	}

	statusMap := make(map[uuid.UUID]string)
	for _, status := range statusEntities {
		statusMap[status.ID] = status.Name
	}

	orderItemsMap := make(map[uuid.UUID][]sqlc.OrderItem)
	for _, item := range orderItemsEntities {
		orderItemsMap[item.OrderID] = append(orderItemsMap[item.OrderID], item)
	}

	orders := make([]domain.Order, 0, len(orderEntities))
	for _, orderEntity := range orderEntities {
		order, err := r.toOrderDomain(&orderEntity, orderItemsMap[orderEntity.ID], statusMap)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *order)
	}

	return &orders, nil
}

func (r *Order) Count(
	ctx context.Context,
) (*int, error) {
	count, err := r.queries.CountOrders(ctx, sqlc.CountOrdersParams{})
	if err != nil {
		return nil, toDomainError(err)
	}
	result := int(count)
	return &result, nil
}

func (r *Order) Get(
	ctx context.Context,
	params domain.OrderRepositoryGetParam,
) (*domain.Order, error) {
	orderEntity, err := r.queries.GetOrder(ctx, sqlc.GetOrderParams{
		ID: params.OrderID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	orderItemsEntities, err := r.queries.ListOrderItems(ctx, sqlc.ListOrderItemsParams{
		OrderID: orderEntity.ID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	statusEntity, err := r.queries.GetOrderStatus(ctx, sqlc.GetOrderStatusParams{
		ID: orderEntity.StatusID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	providerEntity, err := r.queries.GetOrderProvider(ctx, sqlc.GetOrderProviderParams{
		ID: orderEntity.ProviderID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	statusMap := map[uuid.UUID]string{
		statusEntity.ID: statusEntity.Name,
	}

	order, err := r.toOrderDomain(&orderEntity, orderItemsEntities, statusMap)
	if err != nil {
		return nil, err
	}
	order.Provider = domain.OrderProvider(providerEntity.Name)

	return order, nil
}

func (r *Order) Save(
	ctx context.Context,
	params domain.OrderRepositorySaveParam,
) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return toDomainError(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	qtx := r.queries.WithTx(tx)

	statusEntity, err := qtx.GetOrderStatus(ctx, sqlc.GetOrderStatusParams{
		ID:   uuid.Nil,
		Name: string(params.Order.Status),
	})
	if err != nil {
		return toDomainError(err)
	}

	providerEntity, err := qtx.GetOrderProvider(ctx, sqlc.GetOrderProviderParams{
		ID:   uuid.Nil,
		Name: string(params.Order.Provider),
	})
	if err != nil {
		return toDomainError(err)
	}

	err = qtx.UpsertOrder(ctx, sqlc.UpsertOrderParams{
		ID:          params.Order.ID,
		UserID:      params.Order.UserID,
		Email:       params.Order.Email,
		FirstName:   params.Order.FirstName,
		LastName:    params.Order.LastName,
		Address:     params.Order.Address,
		City:        params.Order.City,
		TotalAmount: int64ToNumeric(params.Order.TotalAmount),
		IsPaid:      params.Order.IsPaid,
		ProviderID:  providerEntity.ID,
		StatusID:    statusEntity.ID,
		CreatedAt: pgtype.Timestamptz{
			Time:  params.Order.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  params.Order.UpdatedAt,
			Valid: true,
		},
	})
	if err != nil {
		return toDomainError(err)
	}

	err = r.mergeOrderItems(ctx, qtx, &params.Order)
	if err != nil {
		return toDomainError(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return toDomainError(err)
	}

	return nil
}

func (r *Order) mergeOrderItems(
	ctx context.Context,
	qtx *sqlc.Queries,
	order *domain.Order,
) error {
	if err := qtx.CreateTempTableOrderItems(ctx); err != nil {
		return err
	}

	param := make([]sqlc.InsertTempTableOrderItemsParams, 0, len(order.Items))
	for _, item := range order.Items {
		param = append(param, sqlc.InsertTempTableOrderItemsParams{
			ID:       item.ID,
			Quantity: int32(item.Quantity),
			OrderID:  order.ID,
			Price:    int64ToNumeric(item.Price),
			BookID:   item.BookID,
		})
	}

	if _, err := qtx.InsertTempTableOrderItems(ctx, param); err != nil {
		return err
	}

	if err := qtx.MergeOrderItemsFromTemp(ctx); err != nil {
		return err
	}

	return nil
}

func (r *Order) toOrderDomain(
	orderEntity *sqlc.Order,
	orderItemsEntities []sqlc.OrderItem,
	statusMap map[uuid.UUID]string,
) (*domain.Order, error) {
	items := make([]domain.OrderItem, 0, len(orderItemsEntities))
	for _, itemEntity := range orderItemsEntities {
		items = append(items, domain.OrderItem{
			ID:       itemEntity.ID,
			BookID:   itemEntity.BookID,
			Quantity: int(itemEntity.Quantity),
			Price:    numericToInt64(itemEntity.Price),
		})
	}

	statusName := statusMap[orderEntity.StatusID]

	return &domain.Order{
		ID:          orderEntity.ID,
		UserID:      orderEntity.UserID,
		Address:     orderEntity.Address,
		City:        orderEntity.City,
		Status:      domain.OrderStatus(statusName),
		IsPaid:      orderEntity.IsPaid,
		CreatedAt:   orderEntity.CreatedAt.Time,
		UpdatedAt:   orderEntity.UpdatedAt.Time,
		Email:       orderEntity.Email,
		FirstName:   orderEntity.FirstName,
		LastName:    orderEntity.LastName,
		Items:       items,
		TotalAmount: numericToInt64(orderEntity.TotalAmount),
	}, nil
}
