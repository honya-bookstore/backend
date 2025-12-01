package http

import "time"

type MediaResponseDTO struct {
	ID        string     `json:"id"        binding:"required"`
	URL       string     `json:"url"       binding:"required,url"`
	AltText   string     `json:"altText"   binding:"omitempty,lte=200"`
	Order     int        `json:"order"     binding:"required,gte=0"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt" binding:"omitempty,gtefield=CreatedAt"`
}
