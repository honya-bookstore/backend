package domain

type BookService interface {
	Validate(
		book Book,
	) error
}
