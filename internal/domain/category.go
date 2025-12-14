package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `validate:"required"`
	Slug        string    `validate:"required,gte=2,lte=100"`
	Name        string    `validate:"required,gte=2,lte=100"`
	Description string    `validate:"omitempty,lte=500"`
	CreatedAt   time.Time `validate:"required"`
	UpdatedAt   time.Time `validate:"required,gtefield=CreatedAt"`
	DeletedAt   time.Time `validate:"omitempty,gtefield=CreatedAt"`
}

func NewCategory(slug, name, description string) (*Category, error) {
	now := time.Now()
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	category := &Category{
		ID:          id,
		Name:        name,
		Slug:        slug + id.String(),
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return category, nil
}

func (c *Category) Update(name, description, slug string) {
	if c == nil {
		return
	}
	updated := false
	if name != "" && name != c.Name {
		c.Name = name
		updated = true
	}
	if description != "" && description != c.Description {
		c.Description = description
		updated = true
	}
	if slug != "" && slug != c.Slug {
		c.Slug = slug
		updated = true
	}
	if updated {
		c.UpdatedAt = time.Now()
	}
}

func (c *Category) Remove() {
	now := time.Now()
	c.DeletedAt = now
	c.UpdatedAt = now
}
