package http

import (
	"context"

	"github.com/google/uuid"
)

type MediaApplication interface {
	List(ctx context.Context, param ListMediaRequestDTO) (*PaginationResponseDTO[MediaResponseDTO], error)
	Get(ctx context.Context, param GetMediaRequestDTO) (*MediaResponseDTO, error)
	Create(ctx context.Context, param CreateMediaRequestDTO) (*MediaResponseDTO, error)
	Delete(ctx context.Context, param DeleteMediaRequestDTO) error
	GetUploadImageURL(ctx context.Context) (*UploadImageURLResponseDTO, error)
	GetDeleteImageURL(ctx context.Context, imageID uuid.UUID) (*DeleteImageURLResponseDTO, error)
}
