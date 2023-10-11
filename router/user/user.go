package user

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func UserRouter() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/v1/user")
	v1.GET("/index", func(c *gin.Context) {
		service.GetIndex(c)
	})
	v1.POST("/create", func(c *gin.Context) {
		service.CreateUser(c)
	})

	return r
}
