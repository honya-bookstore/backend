package http

import (
	"github.com/gin-gonic/gin"
)

type LoggingMiddleware interface {
	Handler() gin.HandlerFunc
}
