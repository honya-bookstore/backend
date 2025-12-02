package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type Cart struct {
	cartRepo    domain.CartRepository
	cartService domain.CartService
	bookRepo    domain.BookRepository
	bookService domain.BookService
}

func ProvideCart(
	cartRepo domain.CartRepository,
	cartService domain.CartService,
	bookRepo domain.BookRepository,
	bookService domain.BookService,
) *Cart {
	return &Cart{
		cartRepo:    cartRepo,
		cartService: cartService,
		bookRepo:    bookRepo,
		bookService: bookService,
	}
}

var _ http.CartApplication = &Cart{}

func (c *Cart) Get(ctx context.Context, param http.GetCartRequestDTO) (*http.CartResponseDTO, error) {
	cart, err := c.cartRepo.Get(ctx, domain.CartRepositoryGetParam{
		CartID: param.PathParams.CartID,
	})
	if err != nil {
		return nil, err
	}

	return c.enrichCart(ctx, cart)
}

func (c *Cart) GetByUser(ctx context.Context, param http.GetCartByUserRequestDTO) (*http.CartResponseDTO, error) {
	cart, err := c.cartRepo.Get(ctx, domain.CartRepositoryGetParam{
		UserID: param.PathParams.UserID,
	})
	if err != nil {
		return nil, err
	}

	return c.enrichCart(ctx, cart)
}

func (c *Cart) Create(ctx context.Context, param http.CreateCartRequestDTO) (*http.CartResponseDTO, error) {
	cart, err := domain.NewCart(param.Data.UserID)
	if err != nil {
		return nil, err
	}

	if err := c.cartService.Validate(*cart); err != nil {
		return nil, err
	}

	err = c.cartRepo.Save(ctx, domain.CartRepositorySaveParam{
		Cart: *cart,
	})
	if err != nil {
		return nil, err
	}

	return c.enrichCart(ctx, cart)
}

func (c *Cart) Update(ctx context.Context, param http.UpdateCartItemRequestDTO) (*http.CartResponseDTO, error) {
	cart, err := c.cartRepo.Get(ctx, domain.CartRepositoryGetParam{
		CartID: param.PathParams.CartID,
	})
	if err != nil {
		return nil, err
	}

	cart.UpdateItem(
		param.PathParams.CartItemID,
		param.Data.Quantity,
	)

	if err := c.cartService.Validate(*cart); err != nil {
		return nil, err
	}

	err = c.cartRepo.Save(ctx, domain.CartRepositorySaveParam{
		Cart: *cart,
	})
	if err != nil {
		return nil, err
	}

	return c.enrichCart(ctx, cart)
}

func (c *Cart) AddItem(ctx context.Context, param http.CreateCartItemRequestDTO) (*http.CartResponseDTO, error) {
	// Get or create cart for user
	cart, err := c.cartRepo.Get(ctx, domain.CartRepositoryGetParam{
		UserID: param.Data.UserID,
	})
	if err != nil {
		// If cart doesn't exist, create new one
		cart, err = domain.NewCart(param.Data.UserID)
		if err != nil {
			return nil, err
		}
	}

	// Validate book exists
	book, err := c.bookRepo.Get(ctx, domain.BookRepositoryGetParam{
		BookID: param.Data.BookID,
	})
	if err != nil {
		return nil, err
	}

	// Check stock availability
	if book.StockQuantity < param.Data.Quantity {
		return nil, domain.ErrInvalid
	}

	cartItem, err := domain.NewCartItem(
		param.Data.BookID,
		param.Data.Quantity,
	)
	if err != nil {
		return nil, err
	}

	cart.UpsertItem(*cartItem)

	if err := c.cartService.Validate(*cart); err != nil {
		return nil, err
	}

	err = c.cartRepo.Save(ctx, domain.CartRepositorySaveParam{
		Cart: *cart,
	})
	if err != nil {
		return nil, err
	}

	return c.enrichCart(ctx, cart)
}

func (c *Cart) UpdateItem(ctx context.Context, param http.UpdateCartItemRequestDTO) (*http.CartResponseDTO, error) {
	cart, err := c.cartRepo.Get(ctx, domain.CartRepositoryGetParam{
		CartID: param.PathParams.CartID,
	})
	if err != nil {
		return nil, err
	}

	cart.UpdateItem(
		param.PathParams.CartItemID,
		param.Data.Quantity,
	)

	if err := c.cartService.Validate(*cart); err != nil {
		return nil, err
	}

	err = c.cartRepo.Save(ctx, domain.CartRepositorySaveParam{
		Cart: *cart,
	})
	if err != nil {
		return nil, err
	}

	return c.enrichCart(ctx, cart)
}

func (c *Cart) DeleteItem(ctx context.Context, param http.DeleteCartItemRequestDTO) error {
	cart, err := c.cartRepo.Get(ctx, domain.CartRepositoryGetParam{
		CartID: param.PathParams.CartID,
	})
	if err != nil {
		return err
	}

	cart.RemoveItem(param.PathParams.CartItemID)

	if err := c.cartService.Validate(*cart); err != nil {
		return err
	}

	err = c.cartRepo.Save(ctx, domain.CartRepositorySaveParam{
		Cart: *cart,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Cart) enrichCart(ctx context.Context, cart *domain.Cart) (*http.CartResponseDTO, error) {
	if len(cart.Items) == 0 {
		return http.ToCartResponseDTO(cart, nil), nil
	}

	// Collect unique book IDs
	bookIDs := make([]uuid.UUID, 0, len(cart.Items))
	for _, item := range cart.Items {
		bookIDs = append(bookIDs, item.BookID)
	}

	// Fetch all books at once
	books, err := c.bookRepo.List(
		ctx,
		domain.BookRepositoryListParam{
			BookIDs: bookIDs,
		},
	)
	if err != nil {
		return nil, err
	}

	// Create map for quick lookup
	bookMap := make(map[uuid.UUID]*domain.Book)
	for i := range *books {
		bookMap[(*books)[i].ID] = &(*books)[i]
	}

	return http.ToCartResponseDTO(cart, bookMap), nil
}
