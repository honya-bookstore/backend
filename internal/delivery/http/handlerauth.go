package http

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Handler() gin.HandlerFunc
}
