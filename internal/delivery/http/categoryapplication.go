package http

import "context"

type CategoryApplication interface {
	Create(ctx context.Context, param CreateCategoryRequestDTO) (*CategoryResponseDTO, error)
	List(ctx context.Context, param ListCategoryRequestDTO) (*PaginationResponseDTO[CategoryResponseDTO], error)
	GetBySlug(ctx context.Context, param GetCategoryBySlugRequestDTO) (*CategoryResponseDTO, error)
	Update(ctx context.Context, param UpdateCategoryRequestDTO) (*CategoryResponseDTO, error)
	Delete(ctx context.Context, param DeleteCategoryRequestDTO) error
}
