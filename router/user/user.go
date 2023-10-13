package user

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func UserRouter() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/v1/user")

	v1.POST("/register", func(c *gin.Context) {
		service.UserRegister(c)
	})
	v1.POST("/login", func(c *gin.Context) {
		service.UserLogin(c)
	})

	return r
}
