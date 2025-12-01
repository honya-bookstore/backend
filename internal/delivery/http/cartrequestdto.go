package http

import "github.com/google/uuid"

type GetCartRequestDTO struct {
	PathParams *GetCartRequestPathParams
}

type GetCartRequestPathParams struct {
	CartID uuid.UUID `json:"id" binding:"required" format:"uuid"`
}

type GetCartByUserRequestDTO struct {
	PathParams *GetCartByUserRequestPathParams
}

type GetCartByUserRequestPathParams struct {
	UserID uuid.UUID `json:"user_id" binding:"required" format:"uuid"`
}

type CreateCartRequestDTO struct {
	Data *CreateCartData
}

type CreateCartData struct {
	UserID uuid.UUID `json:"userId" binding:"required"`
}

type CreateCartItemRequestDTO struct {
	Data *CreateCartItemData
}

type CreateCartItemData struct {
	BookID   uuid.UUID `json:"bookId"   binding:"required"`
	Quantity int       `json:"quantity" binding:"required,gt=0"`
	UserID   uuid.UUID `json:"userId"   binding:"required"`
}

type UpdateCartItemRequestDTO struct {
	PathParams *UpdateCartItemRequestPathParams
	Data       *UpdateCartItemData
}

type UpdateCartItemRequestPathParams struct {
	CartID     uuid.UUID `json:"id"      binding:"required" format:"uuid"`
	CartItemID uuid.UUID `json:"item_id" binding:"required" format:"uuid"`
}

type UpdateCartItemData struct {
	Quantity int `json:"quantity"`
}

type DeleteCartItemRequestDTO struct {
	PathParams *DeleteCartItemRequestPathParams
}

type DeleteCartItemRequestPathParams struct {
	CartID     uuid.UUID `json:"id"      binding:"required" format:"uuid"`
	CartItemID uuid.UUID `json:"item_id" binding:"required" format:"uuid"`
}
