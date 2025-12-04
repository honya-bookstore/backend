package http

import (
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	List(ctx *gin.Context)
	GetBySlug(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
