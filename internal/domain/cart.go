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

type CartItem struct {
	ID       uuid.UUID `json:"id"       binding:"required" validate:"required,uuid"`
	BookID   uuid.UUID `json:"bookId"`
	Quantity int       `json:"quantity" binding:"required" validate:"required,gt=0,lte=100"`
}

func NewCart(userID uuid.UUID) (*Cart, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	cart := &Cart{
		ID:        id,
		UserID:    userID,
		Items:     []CartItem{},
		UpdatedAt: time.Now(),
	}
	return cart, nil
}

func NewCartItem(
	BookID uuid.UUID,
	quantity int,
) (*CartItem, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	cartItem := &CartItem{
		ID:       id,
		BookID:   BookID,
		Quantity: quantity,
	}
	return cartItem, nil
}

func (c *Cart) UpsertItem(item CartItem) CartItem {
	for i := range c.Items {
		existingItem := &c.Items[i]
		if existingItem.BookID == item.BookID {
			existingItem.Quantity += item.Quantity
			return *existingItem
		}
	}
	c.Items = append(c.Items, item)
	return item
}

func (c *Cart) RemoveItem(itemID uuid.UUID) {
	filteredItems := []CartItem{}
	for _, item := range c.Items {
		if item.ID != itemID {
			filteredItems = append(filteredItems, item)
		}
	}
	c.Items = filteredItems
}

func (c *Cart) UpdateItem(
	itemID uuid.UUID,
	quantity int,
) {
	for i, item := range c.Items {
		if item.ID == itemID {
			if quantity == 0 {
				c.RemoveItem(itemID)
			} else {
				c.Items[i].Quantity = quantity
			}
			break
		}
	}
}
