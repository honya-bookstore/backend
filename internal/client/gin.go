package client

import "github.com/gin-gonic/gin"

func NewGin() *gin.Engine {
	return gin.New()
}
