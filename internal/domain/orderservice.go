package domain

type OrderService interface {
	Validate(
		media Order,
	) error
}
