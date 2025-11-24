package domain

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID  `json:"id"            binding:"required"                    validate:"required"`
	Title         string     `json:"title"         binding:"required"                    validate:"required,gte=1,lte=200"`
	Description   string     `json:"description"                                         validate:"omitempty,lte=2000"`
	Author        string     `json:"author"                                              validate:"omitempty,lte=200"`
	Price         int64      `json:"price"         binding:"required"                    validate:"required,gt=0"`
	PagesCount    int        `json:"pagesCount"                                          validate:"omitempty,gt=0,lte=10000"`
	YearPublished int        `json:"yearPublished"                                       validate:"omitempty,gte=1000,lte=9999"`
	Publisher     string     `json:"publisher"                                           validate:"omitempty,lte=200"`
	Weight        float64    `json:"weight"                                              validate:"omitempty,gt=0"`
	StockQuantity int        `json:"stockQuantity" binding:"required"                    validate:"required,gte=0"`
	PurchaseCount int        `json:"purchaseCount" binding:"required"                    validate:"required,gte=0"`
	Rating        float32    `json:"rating"        binding:"required"                    validate:"required,gte=0,lte=5"`
	Category      *Category  `json:"category"`
	Media         []Media    `json:"media"         validate:"omitempty,dive"`
	CreatedAt     time.Time  `json:"createdAt"     binding:"required"                    validate:"required"`
	UpdatedAt     time.Time  `json:"updatedAt"     binding:"required"                    validate:"required,gtefield=CreatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"     validate:"omitnil,gtefield=CreatedAt"`
}

func (b *Book) AddMedia(media ...Media) {
	if b.Media == nil {
		b.Media = []Media{}
	}
	b.Media = append(b.Media, media...)
}
