package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/index", func(c *gin.Context) {
		service.GetIndex(c)
	})

	r.POST("/user/create", func(c *gin.Context) {
		service.CreateUser(c)
	})
	return r
}
