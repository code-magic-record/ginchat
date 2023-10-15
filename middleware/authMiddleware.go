package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "未登录",
			})
			c.Abort()
			return
		}
		fmt.Println(cookie, "cookie")
		c.Next()
	}
}
