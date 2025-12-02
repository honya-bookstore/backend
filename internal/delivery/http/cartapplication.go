package http

import "context"

type CartApplication interface {
	Get(ctx context.Context, param GetCartRequestDTO) (*CartResponseDTO, error)
	GetByUser(ctx context.Context, param GetCartByUserRequestDTO) (*CartResponseDTO, error)
	GetMine(ctx context.Context, param GetCartByUserRequestDTO) (*CartResponseDTO, error)
	Create(ctx context.Context, param CreateCartRequestDTO) (*CartResponseDTO, error)
	Update(ctx context.Context, param UpdateCartItemRequestDTO) (*CartResponseDTO, error)
	AddItem(ctx context.Context, param CreateCartItemRequestDTO) (*CartResponseDTO, error)
	UpdateItem(ctx context.Context, param UpdateCartItemRequestDTO) (*CartResponseDTO, error)
	DeleteItem(ctx context.Context, param DeleteCartItemRequestDTO) error
}
