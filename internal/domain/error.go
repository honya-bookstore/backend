package domain

import "errors"

var (
	ErrConflict       = errors.New("conflict")
	ErrExists         = errors.New("already exists")
	ErrForbidden      = errors.New("forbidden")
	ErrInternal       = errors.New("internal error")
	ErrInvalid        = errors.New("invalid data")
	ErrNotFound       = errors.New("not found")
	ErrNotImplemented = errors.New("not implemented")
	ErrServiceError   = errors.New("service error")
	ErrTimeout        = errors.New("timeout")
	ErrUnavailable    = errors.New("service unavailable")
	ErrUnknown        = errors.New("unknown error")
)
