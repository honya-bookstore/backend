package http

import "context"

type OrderApplication interface {
	List(ctx context.Context, param ListOrderRequestDTO) ([]OrderResponseDTO, error)
	Get(ctx context.Context, param GetOrderRequestDTO) (*OrderResponseDTO, error)
	Create(ctx context.Context, param CreateOrderRequestDTO) (*OrderResponseDTO, error)
	Update(ctx context.Context, param UpdateOrderRequestDTO) (*OrderResponseDTO, error)
}
