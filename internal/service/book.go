package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
)

type Book struct {
	validate *validator.Validate
}

func ProvideBook(validate *validator.Validate) *Book {
	return &Book{validate: validate}
}

var _ domain.BookService = (*Book)(nil)

func (b *Book) Validate(book domain.Book) error {
	if err := b.validate.Struct(book); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
