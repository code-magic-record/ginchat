package router

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		service.GetIndex(c)
	})

	return r
}
