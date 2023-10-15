package router

import (
	"ginchat/docs"
	"ginchat/router/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.New()
	docs.SwaggerInfo.BasePath = ""

	// 使用重启和日志中间件
	r.Use(gin.Recovery(), gin.Logger())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 将不同类型的路由分别放在不同的文件夹中
	r.Any("/v1/user/*any", gin.WrapH(user.UserRouter()))

	return r
}
