package http

import (
	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}
