package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

func GoodsAddCategory(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")

	if name == "" {
		c.JSON(400, gin.H{
			"message": "参数异常",
			"code":    0,
		})
		return
	}

	category := models.CategoryBasic{
		Name:     name,
		Describe: desc,
	}

	err := models.CreateCategory(category)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"code":    0,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "添加成功",
		"code":    1,
	})
}
