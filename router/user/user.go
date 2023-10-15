package user

import (
	"ginchat/middleware"
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

	v1.POST("/logout", middleware.AuthMiddleWare(), func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "退出登录成功",
		})

	})

	v1.GET("/getInfo", middleware.AuthMiddleWare(), func(c *gin.Context) {
		service.GetUserInfo(c)
	})
	return r
}
