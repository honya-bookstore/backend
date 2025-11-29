package domain

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID        uuid.UUID `json:"id"        binding:"required"                    validate:"required"`
	URL       string    `json:"url"       binding:"required"                    validate:"required,url"`
	AltText   string    `json:"altText"   validate:"omitempty,lte=200"`
	Order     int       `json:"order"     binding:"required"                    validate:"required,gte=0"`
	BookID    uuid.UUID `json:"bookId"    validate:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"                    validate:"required"`
	DeletedAt time.Time `json:"deletedAt" validate:"omitnil,gtefield=CreatedAt"`
}

func NewMedia(
	url string,
	altText string,
	order int,
	bookID uuid.UUID,
) (*Media, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &Media{
		ID:        id,
		URL:       url,
		AltText:   altText,
		Order:     order,
		BookID:    bookID,
		CreatedAt: now,
	}, nil
}
