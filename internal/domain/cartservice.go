package domain

type CartService interface {
	Validate(
		cart Cart,
	) error
}
