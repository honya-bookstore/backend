package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID  `json:"id"          binding:"required"                    validate:"required"`
	Slug        string     `json:"slug"        binding:"required"                    validate:"required,gte=2,lte=100"`
	Name        string     `json:"name"        binding:"required"                    validate:"required,gte=2,lte=100"`
	Description string     `json:"description" validate:"omitempty,lte=500"`
	CreatedAt   time.Time  `json:"createdAt"   binding:"required"                    validate:"required"`
	UpdatedAt   time.Time  `json:"updatedAt"   binding:"required"                    validate:"required,gtefield=CreatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"   validate:"omitnil,gtefield=CreatedAt"`
}

func NewCategory(slug string, name string, description string) (*Category, error) {
	now := time.Now()
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	category := &Category{
		ID:          id,
		Name:        name,
		Slug:        slug,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return category, nil
}

func (c *Category) Update(name *string, description *string, slug *string) {
	if c == nil {
		return
	}
	updated := false
	if name != nil {
		c.Name = *name
		updated = true
	}
	if description != nil {
		c.Description = *description
		updated = true
	}
	if slug != nil {
		c.Slug = *slug
		updated = true
	}
	if updated {
		c.UpdatedAt = time.Now()
	}
}
