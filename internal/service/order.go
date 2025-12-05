package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
)

type Order struct {
	validate *validator.Validate
}

func ProvideOrder(validate *validator.Validate) *Order {
	return &Order{validate: validate}
}

var _ domain.OrderService = (*Order)(nil)

func (o *Order) Validate(order domain.Order) error {
	if err := o.validate.Struct(order); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
