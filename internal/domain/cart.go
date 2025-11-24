package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `json:"id"        binding:"required"        validate:"required"`
	Items     []CartItem `json:"items"     validate:"omitempty,dive"`
	UserID    uuid.UUID  `json:"userId"    binding:"required"        validate:"required,uuid"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"        validate:"required"`
}

func (c *Cart) AddItems(items ...CartItem) {
	if c.Items == nil {
		c.Items = []CartItem{}
	}
	c.Items = append(c.Items, items...)
}

type CartItem struct {
	ID       uuid.UUID `json:"id"       binding:"required" validate:"required,uuid"`
	Book     *Book     `json:"book"`
	Quantity int       `json:"quantity" binding:"required" validate:"required,gt=0,lte=100"`
}
