package http

import (
	"github.com/gin-gonic/gin"
)

type MetricMiddleware interface {
	Handler() gin.HandlerFunc
}
