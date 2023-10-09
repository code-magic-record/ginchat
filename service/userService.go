package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// CreateUser 创建用户
// @tags 用户管理
// @Summary 创建用户
// @Description 创建用户
// @Accept  json
// @Produce  json
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Param repasswrod query string true "确认密码"
// @Success 200 {string} json "{"message": "创建成功"}"
// @Router /user/create [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repasswrod := c.Query("repasswrod")

	if password != repasswrod {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
	}

	user.Password = password
	// fmt.Println(user)
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "创建成功",
	})
}
