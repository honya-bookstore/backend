package http

import (
	"net/http"

	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type OrderHandlerImpl struct {
	orderApp           OrderApplication
	ErrRequiredOrderID string
	ErrInvalidOrderID  string
}

var _ OrderHandler = &OrderHandlerImpl{}

func ProvideOrderHandler(orderApp OrderApplication) *OrderHandlerImpl {
	return &OrderHandlerImpl{
		orderApp:           orderApp,
		ErrRequiredOrderID: "order_id is required",
		ErrInvalidOrderID:  "invalid order_id",
	}
}

// ListOrders godoc
//
//	@Summary		List all orders
//	@Description	Get all orders
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param	queryParams query ListOrderRequestQueryParams true "Query parameters"
//	@Success		200			{array}		OrderResponseDTO
//	@Failure		500			{object}	Error
//	@Router			/orders [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *OrderHandlerImpl) List(ctx *gin.Context) {
	paginateDto, err := createPaginationRequestDtoFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}
	statusQuery := ctx.Query("status")
	status := domain.OrderStatus(statusQuery)
	orders, err := h.orderApp.List(ctx.Request.Context(), ListOrderRequestDTO{
		QueryParams: &ListOrderRequestQueryParams{
			PaginationRequestDTO: *paginateDto,
			Status:               status,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

// GetOrder godoc
//
//	@Summary		Get order by ID
//	@Description	Get order details by ID
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			pathParams 	path	GetOrderRequestPathParams	true	"Path parameters"
//	@Success		200			{object}	OrderResponseDTO
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/orders/{id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *OrderHandlerImpl) Get(ctx *gin.Context) {
	orderID, ok := pathToUUID(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidOrderID))
		return
	}
	order, err := h.orderApp.Get(ctx.Request.Context(), GetOrderRequestDTO{
		PathParams: &GetOrderRequestPathParams{
			OrderID: orderID,
		},
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, order)
}

// CreateOrder godoc
//
//	@Summary		Create a new order
//	@Description	Create a new order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order	body		CreateOrderData	true	"Order request"
//	@Success		201		{object}	OrderResponseDTO
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/orders [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *OrderHandlerImpl) Create(ctx *gin.Context) {
	var data CreateOrderData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	order, err := h.orderApp.Create(ctx, CreateOrderRequestDTO{
		Data: &data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, order)
}

// UpdateOrder godoc
//
//	@Summary		Update order
//	@Description	Update  order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			pathParams	path		UpdateOrderRequestPathParams	true	"Path parameters"
//	@Param			data body		UpdateOrderData	true	"Update order status request"
//	@Success		200			{object}	OrderResponseDTO
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/orders/{id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *OrderHandlerImpl) Update(ctx *gin.Context) {
	orderID, ok := pathToUUID(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidOrderID))
		return
	}
	var data UpdateOrderData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}
	order, err := h.orderApp.Update(ctx.Request.Context(), UpdateOrderRequestDTO{
		PathParams: &UpdateOrderRequestPathParams{
			OrderID: orderID,
		},
		Data: data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, order)
}
