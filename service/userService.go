package service

import (
	"ginchat/models"
	"ginchat/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	user := models.UserBasic{}
	name := c.PostForm("name")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	phone := c.PostForm("phone")

	hasPhone := models.SearchPhone(phone)
	if hasPhone {
		c.JSON(400, gin.H{
			"message": "手机号已注册",
			"code":    0,
		})
		return
	}
	if password != repassword {
		c.JSON(400, gin.H{
			"message": "两次密码不一致",
			"code":    0,
		})
		return
	}

	user.Name = name
	user.Password = utils.EnCodeMD5(password)
	user.Phone = phone

	err := models.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"code":    0,
		})
		return

	}
	c.JSON(200, gin.H{
		"message": "创建成功",
		"code":    1,
	})
}
