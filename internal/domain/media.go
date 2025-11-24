package domain

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID        uuid.UUID  `json:"id"        binding:"required"                    validate:"required"`
	URL       string     `json:"url"       binding:"required"                    validate:"required,url"`
	AltText   string     `json:"altText"                                         validate:"omitempty,lte=200"`
	Order     int        `json:"order"     binding:"required"                    validate:"required,gte=0"`
	BookID    *uuid.UUID `json:"bookId"    validate:"omitempty,uuid"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"                    validate:"required"`
	DeletedAt *time.Time `json:"deletedAt" validate:"omitnil,gtefield=CreatedAt"`
}
