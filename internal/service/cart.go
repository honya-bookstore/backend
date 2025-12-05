package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
)

type Cart struct {
	validate *validator.Validate
}

func ProvideCart(validate *validator.Validate) *Cart {
	return &Cart{validate: validate}
}

var _ domain.CartService = (*Cart)(nil)

func (c *Cart) Validate(cart domain.Cart) error {
	if err := c.validate.Struct(cart); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
