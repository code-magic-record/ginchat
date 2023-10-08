package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	data := make([]*models.UserBasic, 10)

	// data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}
