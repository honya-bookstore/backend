package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type Order struct {
	orderRepo           domain.OrderRepository
	orderService        domain.OrderService
	bookRepo            domain.BookRepository
	bookService         domain.BookService
	cartRepo            domain.CartRepository
	VNPayPaymentService VNPayPaymentService
}

func ProvideOrder(
	orderRepo domain.OrderRepository,
	orderService domain.OrderService,
	bookRepo domain.BookRepository,
	bookService domain.BookService,
	cartRepo domain.CartRepository,
	VNPayPaymentService VNPayPaymentService,
) *Order {
	return &Order{
		orderRepo:           orderRepo,
		orderService:        orderService,
		bookRepo:            bookRepo,
		bookService:         bookService,
		cartRepo:            cartRepo,
		VNPayPaymentService: VNPayPaymentService,
	}
}

var _ http.OrderApplication = &Order{}

func (o *Order) List(ctx context.Context, param http.ListOrderRequestDTO) ([]http.OrderResponseDTO, error) {
	queryParams := param.QueryParams
	if queryParams == nil {
		queryParams = &http.ListOrderRequestQueryParams{}
	}

	orders, err := o.orderRepo.List(
		ctx,
		domain.OrderRepositoryListParam{
			Status: queryParams.Status,
			Limit:  queryParams.Limit,
			Offset: (queryParams.Page - 1) * queryParams.Limit,
		},
	)
	if err != nil {
		return nil, err
	}

	orderDtos := make([]http.OrderResponseDTO, 0, len(*orders))
	for i := range *orders {
		order := &(*orders)[i]
		orderDto, err := o.enrichOrder(ctx, order, "")
		if err != nil {
			return nil, err
		}
		orderDtos = append(orderDtos, *orderDto)
	}

	return orderDtos, nil
}

func (o *Order) Get(ctx context.Context, param http.GetOrderRequestDTO) (*http.OrderResponseDTO, error) {
	order, err := o.orderRepo.Get(ctx, domain.OrderRepositoryGetParam{
		OrderID: param.PathParams.OrderID,
	})
	if err != nil {
		return nil, err
	}

	return o.enrichOrder(ctx, order, "")
}

func (o *Order) Create(ctx context.Context, param http.CreateOrderRequestDTO) (*http.OrderResponseDTO, error) {
	// check out all item in cart
	cart, err := o.cartRepo.Get(ctx, domain.CartRepositoryGetParam{
		UserID: param.Data.UserID,
	})
	if err != nil {
		return nil, err
	}
	if len(cart.Items) <= 0 {
		return nil, domain.ErrInvalid
	}
	bookIDs := make([]uuid.UUID, 0, len(cart.Items))
	for _, item := range cart.Items {
		bookIDs = append(bookIDs, item.BookID)
	}

	books, err := o.bookRepo.List(ctx, domain.BookRepositoryListParam{
		BookIDs: bookIDs,
	})
	if err != nil {
		return nil, err
	}

	bookMap := make(map[uuid.UUID]*domain.Book, len(*books))
	for i := range *books {
		bookMap[(*books)[i].ID] = &(*books)[i]
	}

	items := make([]domain.OrderItem, 0, len(cart.Items))
	for _, itemData := range cart.Items {
		book, ok := bookMap[itemData.BookID]
		if !ok {
			return nil, domain.ErrNotFound
		}

		if book.StockQuantity < itemData.Quantity {
			return nil, domain.ErrInvalid
		}

		item, err := domain.NewOrderItem(
			itemData.BookID,
			itemData.Quantity,
			book.Price,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}

	order, err := domain.NewOrder(
		param.Data.UserID,
		param.Data.Email,
		param.Data.FirstName,
		param.Data.LastName,
		param.Data.Address,
		param.Data.Provider,
		param.Data.City,
		param.Data.Phone,
		items,
	)
	if err != nil {
		return nil, err
	}

	if err := o.orderService.Validate(*order); err != nil {
		return nil, err
	}

	err = o.orderRepo.Save(ctx, domain.OrderRepositorySaveParam{
		Order: *order,
	})
	if err != nil {
		return nil, err
	}
	var paymentURL string
	var paymentServiceErr error
	switch param.Data.Provider {
	case domain.OrderProvider(domain.PaymentProviderCOD):
	case domain.OrderProvider(domain.PaymentProviderVNPAY):
		paymentURL, paymentServiceErr = o.VNPayPaymentService.GetPaymentURL(ctx, GetPaymentURLVNPayParam{
			Order:     order,
			ReturnURL: param.Data.ReturnURL,
		})
	default:
		paymentServiceErr = domain.ErrInvalid
	}

	if paymentServiceErr != nil {
		return nil, paymentServiceErr
	}
	return o.enrichOrder(ctx, order, paymentURL)
}

func (o *Order) Update(ctx context.Context, param http.UpdateOrderRequestDTO) (*http.OrderResponseDTO, error) {
	order, err := o.orderRepo.Get(ctx, domain.OrderRepositoryGetParam{
		OrderID: param.PathParams.OrderID,
	})
	if err != nil {
		return nil, err
	}

	order.Update(
		param.Data.Address,
		param.Data.Status,
		param.Data.IsPaid,
	)

	if err := o.orderService.Validate(*order); err != nil {
		return nil, err
	}

	err = o.orderRepo.Save(ctx, domain.OrderRepositorySaveParam{
		Order: *order,
	})
	if err != nil {
		return nil, err
	}

	return o.enrichOrder(ctx, order, "")
}

func (o *Order) VerifyVNPayIPN(ctx context.Context, param http.VerifyVNPayIPNRequestDTO) (*http.VerifyVNPayIPNResponseDTO, error) {
	verifyParam := VerifyIPNVNPayParam{
		Amount:            param.QueryParams.Amount,
		BankTranNo:        param.QueryParams.BankTranNo,
		BankCode:          param.QueryParams.BankCode,
		CardType:          param.QueryParams.CardType,
		OrderInfo:         param.QueryParams.OrderInfo,
		PayDate:           param.QueryParams.PayDate,
		ResponseCode:      param.QueryParams.ResponseCode,
		SecureHash:        param.QueryParams.SecureHash,
		TmnCode:           param.QueryParams.TmnCode,
		TransactionNo:     param.QueryParams.TransactionNo,
		TransactionStatus: param.QueryParams.TransactionStatus,
		TxnRef:            param.QueryParams.TxnRef,
	}
	rspCode, message, err := o.VNPayPaymentService.VerifyIPN(
		ctx,
		verifyParam,
		o.getOrder,
		o.onVerifySuccess,
		o.onVerifyFailure,
	)
	return &http.VerifyVNPayIPNResponseDTO{
		RspCode: rspCode,
		Message: message,
	}, err
}

func (o *Order) enrichOrder(ctx context.Context, order *domain.Order, returnURL string) (*http.OrderResponseDTO, error) {
	if len(order.Items) == 0 {
		return http.ToOrderResponseDTO(order, nil, returnURL), nil
	}

	bookIDs := make([]uuid.UUID, 0, len(order.Items))
	for _, item := range order.Items {
		bookIDs = append(bookIDs, item.BookID)
	}

	books, err := o.bookRepo.List(
		ctx,
		domain.BookRepositoryListParam{
			BookIDs: bookIDs,
		},
	)
	if err != nil {
		return nil, err
	}

	bookMap := make(map[uuid.UUID]*domain.Book)
	for i := range *books {
		bookMap[(*books)[i].ID] = &(*books)[i]
	}

	return http.ToOrderResponseDTO(order, bookMap, returnURL), nil
}

func (o *Order) getOrder(
	ctx context.Context,
	orderID uuid.UUID,
) (*domain.Order, error) {
	return o.orderRepo.Get(ctx, domain.OrderRepositoryGetParam{
		OrderID: orderID,
	})
}

func (o *Order) onVerifySuccess(
	ctx context.Context,
	order *domain.Order,
) error {
	order.Update(
		order.Address,
		domain.OrderStatusProcessing,
		true,
	)
	cart, err := o.cartRepo.Get(ctx, domain.CartRepositoryGetParam{
		UserID: order.UserID,
	})
	if err != nil {
		return err
	}
	cart.ClearItems()
	err = o.cartRepo.Save(ctx, domain.CartRepositorySaveParam{
		Cart: *cart,
	})
	if err != nil {
		return err
	}
	bookIDs := make([]uuid.UUID, 0, len(order.Items))
	for _, item := range order.Items {
		bookIDs = append(bookIDs, item.BookID)
	}
	books, err := o.bookRepo.List(ctx, domain.BookRepositoryListParam{
		BookIDs: bookIDs,
	})
	if err != nil {
		return err
	}
	bookIDBookMap := make(map[uuid.UUID]*domain.Book)
	for i := range *books {
		book := &(*books)[i]
		bookIDBookMap[book.ID] = book
	}
	for _, item := range order.Items {
		book, ok := bookIDBookMap[item.BookID]
		if !ok {
			return domain.ErrNotFound
		}
		book.DecreaseQuantity(item.Quantity)
	}
	for _, book := range bookIDBookMap {
		err := o.bookService.Validate(*book)
		if err != nil {
			return err
		}
		err = o.bookRepo.Save(ctx, domain.BookRepositorySaveParam{
			Book: *book,
		})
		if err != nil {
			return err
		}
	}
	return o.orderRepo.Save(ctx, domain.OrderRepositorySaveParam{
		Order: *order,
	})
}

func (o *Order) onVerifyFailure(
	ctx context.Context,
	order *domain.Order,
) error {
	order.Update(
		order.Address,
		domain.OrderStatusCancelled,
		false,
	)
	return o.orderRepo.Save(ctx, domain.OrderRepositorySaveParam{
		Order: *order,
	})
}
