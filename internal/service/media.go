package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
)

type Media struct {
	validate *validator.Validate
}

func ProvideMedia(validate *validator.Validate) *Media {
	return &Media{validate: validate}
}

var _ domain.MediaService = (*Media)(nil)

func (m *Media) Validate(media domain.Media) error {
	if err := m.validate.Struct(media); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
