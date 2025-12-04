package http

import (
	"github.com/gin-gonic/gin"
)

type CartHandler interface {
	Get(ctx *gin.Context)
	GetByUser(ctx *gin.Context)
	GetMine(ctx *gin.Context)
	Create(ctx *gin.Context)
	CreateItem(ctx *gin.Context)
	UpdateItem(ctx *gin.Context)
	DeleteItem(ctx *gin.Context)
}
