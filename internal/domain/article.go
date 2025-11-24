package domain

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID          uuid.UUID  `json:"id"          binding:"required"                    validate:"required"`
	Slug        string     `json:"slug"        binding:"required"                    validate:"required,gte=2,lte=200"`
	Title       string     `json:"title"       binding:"required"                    validate:"required,gte=2,lte=200"`
	ThumbnailID *uuid.UUID `json:"thumbnailId" validate:"omitempty,uuid"`
	Thumbnail   *Media     `json:"thumbnail"`
	Content     string     `json:"content"                                           validate:"omitempty,lte=50000"`
	Tags        []string   `json:"tags"        validate:"omitempty,dive,gte=1,lte=50"`
	UserID      uuid.UUID  `json:"userId"      binding:"required"                    validate:"required,uuid"`
	CreatedAt   time.Time  `json:"createdAt"   binding:"required"                    validate:"required"`
	UpdatedAt   time.Time  `json:"updatedAt"   binding:"required"                    validate:"required,gtefield=CreatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"   validate:"omitnil,gtefield=CreatedAt"`
}
