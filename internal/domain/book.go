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
	CategoryID    uuid.UUID   `json:"categoryId"`
	MediaIDs      []uuid.UUID `json:"mediaIds"      validate:"omitempty,dive"`
	CreatedAt     time.Time   `json:"createdAt"     binding:"required"                     validate:"required"`
	UpdatedAt     time.Time   `json:"updatedAt"     binding:"required"                     validate:"required,gtefield=CreatedAt"`
	DeletedAt     *time.Time  `json:"deletedAt"     validate:"omitnil,gtefield=CreatedAt"`
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
	categoryID uuid.UUID,
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
		CategoryID:    categoryID,
		MediaIDs:      mediaIDs,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, err
}

func (b *Book) Update(
	title *string,
	description *string,
	author *string,
	price *int64,
	pagesCount *int,
	yearPublished *int,
	publisher *string,
	weight *float64,
	stockQuantity *int,
	categoryID *uuid.UUID,
) {
	updated := false
	if title != nil {
		b.Title = *title
		updated = true
	}
	if description != nil {
		b.Description = *description
		updated = true
	}
	if author != nil {
		b.Author = *author
		updated = true
	}
	if price != nil {
		b.Price = *price
		updated = true
	}
	if pagesCount != nil {
		b.PagesCount = *pagesCount
		updated = true
	}
	if yearPublished != nil {
		b.YearPublished = *yearPublished
		updated = true
	}
	if publisher != nil {
		b.Publisher = *publisher
		updated = true
	}
	if weight != nil {
		b.Weight = *weight
		updated = true
	}
	if stockQuantity != nil {
		b.StockQuantity = *stockQuantity
		updated = true
	}
	if categoryID != nil {
		b.CategoryID = *categoryID
		updated = true
	}
	if updated {
		b.UpdatedAt = time.Now()
	}
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

func (b *Book) Remove() {
	now := time.Now()
	if b.DeletedAt == nil {
		b.DeletedAt = &now
	}
}
