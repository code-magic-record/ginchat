package foods

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func FoodsRouter() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/v1/foods")
	v1.POST("/add_category", service.GoodsAddCategory)

	return r
}
