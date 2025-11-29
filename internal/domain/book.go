package domain

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID   `json:"id"            binding:"required"                     validate:"required"`
	Title         string      `json:"title"         binding:"required"                     validate:"required,gte=1,lte=200"`
	Description   string      `json:"description"   validate:"omitempty,lte=2000"`
	Author        string      `json:"author"        validate:"omitempty,lte=200"`
	Price         int64       `json:"price"         binding:"required"                     validate:"required,gt=0"`
	PagesCount    int         `json:"pagesCount"    validate:"omitempty,gt=0,lte=10000"`
	YearPublished int         `json:"yearPublished" validate:"omitempty,gte=1000,lte=9999"`
	Publisher     string      `json:"publisher"     validate:"omitempty,lte=200"`
	Weight        float64     `json:"weight"        validate:"omitempty,gt=0"`
	StockQuantity int         `json:"stockQuantity" binding:"required"                     validate:"required,gte=0"`
	PurchaseCount int         `json:"purchaseCount" binding:"required"                     validate:"required,gte=0"`
	Rating        float32     `json:"rating"        binding:"required"                     validate:"required,gte=0,lte=5"`
	CategoryIDs   []uuid.UUID `json:"categoryIds"`
	MediaIDs      []uuid.UUID `json:"mediaIds"      validate:"omitempty,dive"`
	CreatedAt     time.Time   `json:"createdAt"     binding:"required"                     validate:"required"`
	UpdatedAt     time.Time   `json:"updatedAt"     binding:"required"                     validate:"required,gtefield=CreatedAt"`
	DeletedAt     time.Time   `json:"deletedAt"     validate:"omitnil,gtefield=CreatedAt"`
}

func NewBook(
	title string,
	description string,
	author string,
	price int64,
	pagesCount int,
	yearPublished int,
	publisher string,
	weight float64,
	stockQuantity int,
	categoryID []uuid.UUID,
	mediaIDs []uuid.UUID,
) (*Book, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &Book{
		ID:            id,
		Title:         title,
		Description:   description,
		Author:        author,
		Price:         price,
		PagesCount:    pagesCount,
		YearPublished: yearPublished,
		Publisher:     publisher,
		Weight:        weight,
		StockQuantity: stockQuantity,
		PurchaseCount: 0,
		Rating:        0,
		CategoryIDs:   categoryID,
		MediaIDs:      mediaIDs,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, err
}

func (a *Book) AddMediaIDs(newMediaIDs ...uuid.UUID) {
	a.MediaIDs = append(a.MediaIDs, newMediaIDs...)
}

func (a *Book) RemoveMediaIDs(mediaIDsToRemove ...uuid.UUID) {
	mediaIDMap := make(map[uuid.UUID]struct{})
	for _, mediaID := range mediaIDsToRemove {
		mediaIDMap[mediaID] = struct{}{}
	}

	filteredMediaIDs := make([]uuid.UUID, 0, len(a.MediaIDs)-len(mediaIDsToRemove))
	for _, mediaID := range a.MediaIDs {
		if _, found := mediaIDMap[mediaID]; !found {
			filteredMediaIDs = append(filteredMediaIDs, mediaID)
		}
	}
	a.MediaIDs = filteredMediaIDs
}
