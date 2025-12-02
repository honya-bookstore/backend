package domain

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID   `validate:"required"`
	Title         string      `validate:"required,gte=1,lte=200"`
	Description   string      `validate:"omitempty,lte=2000"`
	Author        string      `validate:"required,lte=200"`
	Price         int64       `validate:"required,gt=0"`
	PagesCount    int         `validate:"required,gt=0,lte=10000"`
	YearPublished int         `validate:"required,gte=1000,lte=9999"`
	Publisher     string      `validate:"required,lte=200"`
	Weight        float64     `validate:"omitempty,gt=0"`
	StockQuantity int         `validate:"required,gte=0"`
	PurchaseCount int         `validate:"required,gte=0"`
	Rating        float64     `validate:"required,gte=0,lte=5"`
	CategoryIDs   []uuid.UUID `validate:"required"`
	MediaIDs      []uuid.UUID `validate:"required,dive"`
	CreatedAt     time.Time   `validate:"required"`
	UpdatedAt     time.Time   `validate:"required,gtefield=CreatedAt"`
	DeletedAt     time.Time   `validate:"omitempty,gtefield=CreatedAt"`
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

func (b *Book) Update(
	title string,
	description string,
	author string,
	price int64,
	pagesCount int,
	yearPublished int,
	publisher string,
	weight float64,
	stockQuantity int,
	categoryIDs []uuid.UUID,
) {
	updated := false
	if title != "" && title != b.Title {
		b.Title = title
		updated = true
	}
	if description != "" && description != b.Description {
		b.Description = description
		updated = true
	}
	if author != "" && author != b.Author {
		b.Author = author
		updated = true
	}
	if price > 0 && price != b.Price {
		b.Price = price
		updated = true
	}
	if pagesCount > 0 && pagesCount != b.PagesCount {
		b.PagesCount = pagesCount
		updated = true
	}
	if yearPublished != 0 && yearPublished != b.YearPublished {
		b.YearPublished = yearPublished
		updated = true
	}
	if publisher != "" && publisher != b.Publisher {
		b.Publisher = publisher
		updated = true
	}
	if weight > 0 && weight != b.Weight {
		b.Weight = weight
		updated = true
	}
	if stockQuantity >= 0 && stockQuantity != b.StockQuantity {
		b.StockQuantity = stockQuantity
		updated = true
	}
	if categoryIDs != nil {
		b.CategoryIDs = categoryIDs
		updated = true
	}
	if updated {
		b.UpdatedAt = time.Now()
	}
}

func (b *Book) AddMediaIDs(newMediaIDs ...uuid.UUID) {
	b.MediaIDs = append(b.MediaIDs, newMediaIDs...)
}

func (b *Book) RemoveMediaIDs(mediaIDsToRemove ...uuid.UUID) {
	mediaIDMap := make(map[uuid.UUID]struct{})
	for _, mediaID := range mediaIDsToRemove {
		mediaIDMap[mediaID] = struct{}{}
	}

	filteredMediaIDs := make([]uuid.UUID, 0, len(b.MediaIDs)-len(mediaIDsToRemove))
	for _, mediaID := range b.MediaIDs {
		if _, found := mediaIDMap[mediaID]; !found {
			filteredMediaIDs = append(filteredMediaIDs, mediaID)
		}
	}
	b.MediaIDs = filteredMediaIDs
}

func (b *Book) Remove() {
	b.DeletedAt = time.Now()
}
