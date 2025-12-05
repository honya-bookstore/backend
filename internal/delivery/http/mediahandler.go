package http

import (
	"github.com/gin-gonic/gin"
)

type MediaHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetUploadImageURL(ctx *gin.Context)
	GetDeleteImageURL(ctx *gin.Context)
}
