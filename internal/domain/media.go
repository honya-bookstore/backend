package domain

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID        uuid.UUID `validate:"required"`
	URL       string    `validate:"required,url"`
	AltText   string    `validate:"omitempty,lte=200"`
	CreatedAt time.Time `validate:"required"`
	DeletedAt time.Time `validate:"omitempty,gtefield=CreatedAt"`
}

func NewMedia(
	url string,
	altText string,
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
		CreatedAt: now,
	}, nil
}

func (m *Media) Delete() {
	m.DeletedAt = time.Now()
}
