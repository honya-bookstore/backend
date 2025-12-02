package http

import (
	"time"

	"backend/internal/domain"
)

type MediaResponseDTO struct {
	ID        string     `json:"id"        binding:"required"`
	URL       string     `json:"url"       binding:"required,url"`
	AltText   string     `json:"altText"   binding:"omitempty,lte=200"`
	Order     int        `json:"order"     binding:"required,gte=0"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt" binding:"omitempty,gtefield=CreatedAt"`
}

func ToMediaResponseDTO(media *domain.Media) *MediaResponseDTO {
	if media == nil {
		return nil
	}

	var deletedAt *time.Time
	if !media.DeletedAt.IsZero() {
		deletedAt = &media.DeletedAt
	}

	return &MediaResponseDTO{
		ID:        media.ID.String(),
		URL:       media.URL,
		AltText:   media.AltText,
		Order:     media.Order,
		CreatedAt: media.CreatedAt,
		DeletedAt: deletedAt,
	}
}

func ToMediaResponseDTOList(media []domain.Media) []MediaResponseDTO {
	result := make([]MediaResponseDTO, 0, len(media))
	for _, m := range media {
		dto := ToMediaResponseDTO(&m)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}
