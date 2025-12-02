package http

import "context"

type BookApplication interface {
	List(ctx context.Context, param ListBookRequestDTO) (*PaginationResponseDTO[BookResponseDTO], error)
	Get(ctx context.Context, param GetBookRequestDTO) (*BookResponseDTO, error)
	Create(ctx context.Context, param CreateBookRequestDTO) (*BookResponseDTO, error)
	Update(ctx context.Context, param UpdateBookRequestDTO) (*BookResponseDTO, error)
	Delete(ctx context.Context, param DeleteBookRequestDTO) error
}
