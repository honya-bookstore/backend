package http

import (
	"errors"

	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func NewError(message string) *Error {
	return &Error{Message: message}
}

func SendError(ctx *gin.Context, err error) {
	var httpErrCode int
	switch {
	case errors.Is(err, domain.ErrNotFound):
		httpErrCode = 404
	case errors.Is(err, domain.ErrInvalid):
		httpErrCode = 400
	case errors.Is(err, domain.ErrExists), errors.Is(err, domain.ErrConflict):
		httpErrCode = 409
	case errors.Is(err, domain.ErrForbidden):
		httpErrCode = 403
	case errors.Is(err, domain.ErrInternal):
		httpErrCode = 500
	case errors.Is(err, domain.ErrUnavailable):
		httpErrCode = 503
	case errors.Is(err, domain.ErrTimeout):
		httpErrCode = 504
	default:
		httpErrCode = 500
	}
	ctx.JSON(httpErrCode, NewError(err.Error()))
}
