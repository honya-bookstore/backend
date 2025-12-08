package application

import (
	"context"

	"backend/internal/delivery/http"

	"github.com/google/uuid"
)

type MediaObjectStorage interface {
	GetUploadImageURL(ctx context.Context) (*http.UploadImageURLResponseDTO, error)
	GetDeleteImageURL(ctx context.Context, imageID uuid.UUID) (*http.DeleteImageURLResponseDTO, error)
	PersistImageFromTemp(ctx context.Context, key string, imageID uuid.UUID) error
	BuildMediaURL(mediaID uuid.UUID) string
}
