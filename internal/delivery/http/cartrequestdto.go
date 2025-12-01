package http

import "github.com/google/uuid"

type GetCartRequestDTO struct {
	ID uuid.UUID
}

type GetCartByUserRequestDTO struct {
	UserID uuid.UUID
}

type CreateCartRequestDTO struct {
	Data *CreateCartData
}

type CreateCartData struct {
	UserID uuid.UUID `json:"userId" binding:"required"`
}

type CreateCartItemRequestDTO struct {
	UserID uuid.UUID
	CartID uuid.UUID
	Data   *CreateCartItemData
}

type CreateCartItemData struct {
	BookID   uuid.UUID `json:"bookId"   binding:"required"`
	Quantity int       `json:"quantity" binding:"required,gt=0"`
}

type UpdateCartItemRequestDTO struct {
	UserID     uuid.UUID
	CartID     uuid.UUID
	CartItemID uuid.UUID
	Data       *UpdateCartItemData
}

type UpdateCartItemData struct {
	Quantity int `json:"quantity"`
}

type DeleteCartItemRequestDTO struct {
	UserID     uuid.UUID
	CartID     uuid.UUID
	CartItemID uuid.UUID
}
