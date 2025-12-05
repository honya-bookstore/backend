package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartHandlerImpl struct {
	cartApp               CartApplication
	ErrRequiredCartID     string
	ErrInvalidCartID      string
	ErrRequiredCartItemID string
	ErrInvalidCartItemID  string
	ErrInvalidUserID      string
}

var _ CartHandler = &CartHandlerImpl{}

func ProvideCartHandler(cartApp CartApplication) *CartHandlerImpl {
	return &CartHandlerImpl{
		cartApp:               cartApp,
		ErrRequiredCartID:     "cart_id is required",
		ErrInvalidCartID:      "invalid cart_id",
		ErrRequiredCartItemID: "cart_item_id is required",
		ErrInvalidCartItemID:  "invalid cart_item_id",
		ErrInvalidUserID:      "invalid user_id",
	}
}

// GetCart godoc
//
//	@Summary		Get cart by ID
//	@Description	Get cart details by ID
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		GetCartRequestPathParams	true	"Path parameters"
//	@Success		200			{object}	CartResponseDTO
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/cart/{id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) Get(ctx *gin.Context) {
	cartID, ok := pathToUUID(ctx, "id")
	if cartID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}
	cart, err := h.cartApp.Get(ctx, GetCartRequestDTO{
		PathParams: &GetCartRequestPathParams{
			CartID: cartID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

// GetCartByUser godoc
//
//	@Summary		Get cart by user ID
//	@Description	Get cart details by user ID
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		GetCartByUserRequestPathParams	true	"Path parameters"
//	@Success		200			{object}	CartResponseDTO
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/cart/user/{id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) GetByUser(ctx *gin.Context) {
	userID, ok := pathToUUID(ctx, "user_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidUserID))
		return
	}
	cart, err := h.cartApp.GetByUser(ctx, GetCartByUserRequestDTO{
		PathParams: &GetCartByUserRequestPathParams{
			UserID: userID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

// GetMyCart godoc
//
//	@Summary		Get current user's cart
//	@Description	Get cart details for the authenticated user
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	CartResponseDTO
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/cart/me [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) GetMine(ctx *gin.Context) {
	userID, ok := ctxValueToUUID(ctx, "userID")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidUserID))
		return
	}
	cart, err := h.cartApp.GetByUser(ctx, GetCartByUserRequestDTO{
		PathParams: &GetCartByUserRequestPathParams{
			UserID: userID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

// CreateCart godoc
//
//	@Summary		Create a new cart
//	@Description	Create a new cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			cart	body		CreateCartData	true	"Cart request"
//	@Success		201		{object}	CartResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/cart [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) Create(ctx *gin.Context) {
	var data CreateCartData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	cart, err := h.cartApp.Create(ctx, CreateCartRequestDTO{
		Data: &data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, cart)
}

// CreateCartItem godoc
//
//	@Summary		Add item to cart
//	@Description	Add a new item to the cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		CreateCartItemRequestPathParams	true	"Path parameters"
//	@Param			item		body		CreateCartItemData	true	"Cart item request"
//	@Success		201			{object}	CartResponseDTO
//	@Failure		400			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/cart/{id}/items [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) CreateItem(ctx *gin.Context) {
	cartID, ok := pathToUUID(ctx, "id")
	if cartID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}

	var data CreateCartItemData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	cartItem, err := h.cartApp.CreateItem(ctx, CreateCartItemRequestDTO{
		PathParams: &CreateCartItemRequestPathParams{
			CartID: cartID,
		},
		Data: &data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, cartItem)
}

// UpdateCartItem godoc
//
//	@Summary		Update cart item
//	@Description	Update cart item quantity
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		UpdateCartItemRequestPathParams	true	"Path parameters"
//	@Param			data		body		UpdateCartItemData	true	"Update cart item request"
//	@Success		200			{object}	CartResponseDTO
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/cart/{id}/items/{item_id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) UpdateItem(ctx *gin.Context) {
	cartID, ok := pathToUUID(ctx, "id")
	if cartID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}

	cartItemID, ok := pathToUUID(ctx, "item_id")
	if cartItemID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartItemID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartItemID))
		return
	}

	var data UpdateCartItemData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	cartItem, err := h.cartApp.UpdateItem(ctx, UpdateCartItemRequestDTO{
		PathParams: &UpdateCartItemRequestPathParams{
			CartID:     cartID,
			CartItemID: cartItemID,
		},
		Data: &data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, cartItem)
}

// DeleteCartItem godoc
//
//	@Summary		Delete cart item
//	@Description	Remove item from cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		DeleteCartItemRequestPathParams	true	"Path parameters"
//	@Success		204
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/cart/{id}/items/{item_id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) DeleteItem(ctx *gin.Context) {
	cartID, ok := pathToUUID(ctx, "id")
	if cartID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}

	itemID, ok := pathToUUID(ctx, "item_id")
	if itemID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartItemID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartItemID))
		return
	}

	err := h.cartApp.DeleteItem(ctx, DeleteCartItemRequestDTO{
		PathParams: &DeleteCartItemRequestPathParams{
			CartID:     cartID,
			CartItemID: itemID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
