package ws

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func WsRouter() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/v1/ws")

	v1.GET("/connect", service.HandlerWs)
	return r
}
