package http

import (
	"github.com/gin-gonic/gin"
)

type RoleMiddleware interface {
	Handler([]UserRole) gin.HandlerFunc
}
