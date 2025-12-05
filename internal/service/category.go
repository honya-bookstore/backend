package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
)

type Category struct {
	validate *validator.Validate
}

func ProvideCategory(validate *validator.Validate) *Category {
	return &Category{validate: validate}
}

var _ domain.CategoryService = (*Category)(nil)

func (c *Category) Validate(category domain.Category) error {
	if err := c.validate.Struct(category); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
