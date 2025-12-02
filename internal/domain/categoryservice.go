package domain

type CategoryService interface {
	Validate(
		category Category,
	) error
}
