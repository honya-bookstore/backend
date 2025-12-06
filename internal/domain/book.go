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
	Price         int64       `validate:"gt=0"`
	PagesCount    int         `validate:"gt=0,lte=10000"`
	YearPublished int         `validate:"gte=1000,lte=9999"`
	Publisher     string      `validate:"required,lte=200"`
	Weight        float64     `validate:"gte=0"`
	StockQuantity int         `validate:"gte=0"`
	PurchaseCount int         `validate:"gte=0"`
	Rating        float64     `validate:"gte=0,lte=5"`
	CategoryIDs   []uuid.UUID `validate:"required"`
	Medium        []BookMedia `validate:"required,dive"`
	CreatedAt     time.Time   `validate:"required"`
	UpdatedAt     time.Time   `validate:"required,gtefield=CreatedAt"`
	DeletedAt     time.Time   `validate:"omitempty,gtefield=CreatedAt"`
}

type BookMedia struct {
	MediaID uuid.UUID `validate:"required"`
	IsCover bool
	Order   int `validate:"gte=0"`
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
	bookMedium []BookMedia,
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
		Medium:        bookMedium,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, err
}

func NewBookMedia(
	MediaID uuid.UUID,
	isCover bool,
	Order int,
) *BookMedia {
	return &BookMedia{
		MediaID: MediaID,
		IsCover: isCover,
		Order:   Order,
	}
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
	media []BookMedia,
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
	if media != nil {
		b.Medium = media
		updated = true
	}
	if updated {
		b.UpdatedAt = time.Now()
	}
}

func (b *Book) Remove() {
	b.DeletedAt = time.Now()
}

func (b *Book) DecreaseQuantity(quantity int) {
	if quantity <= 0 {
		return
	}
	if b.StockQuantity >= quantity {
		b.StockQuantity -= quantity
	} else {
		b.StockQuantity = 0
	}
	b.PurchaseCount += quantity
	b.UpdatedAt = time.Now()
}
