package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type CategoryResponseDTO struct {
	ID          uuid.UUID  `json:"id"          binding:"required"`
	Slug        string     `json:"slug"        binding:"required"`
	Name        string     `json:"name"        binding:"required"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"   binding:"required"`
	UpdatedAt   time.Time  `json:"updatedAt"   binding:"required"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

func ToCategoryResponseDTO(cat *domain.Category) *CategoryResponseDTO {
	if cat == nil {
		return nil
	}

	var deletedAt *time.Time
	if !cat.DeletedAt.IsZero() {
		deletedAt = &cat.DeletedAt
	}

	return &CategoryResponseDTO{
		ID:          cat.ID,
		Slug:        cat.Slug,
		Name:        cat.Name,
		Description: cat.Description,
		CreatedAt:   cat.CreatedAt,
		UpdatedAt:   cat.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

func ToCategoryResponseDTOList(cats []domain.Category) []CategoryResponseDTO {
	result := make([]CategoryResponseDTO, 0, len(cats))
	for _, cat := range cats {
		dto := ToCategoryResponseDTO(&cat)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}
