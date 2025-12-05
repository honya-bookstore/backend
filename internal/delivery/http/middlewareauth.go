package http

import (
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	Handler() gin.HandlerFunc
}
