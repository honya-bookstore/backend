package http

import (
	"time"

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
