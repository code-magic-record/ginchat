package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// ListAccounts godoc
// @Tags         示例
// @Accept       json
// @Produce      json
// @Success      200  {array}   string "OK"
// @Router       /index [get]
func GetIndex(c *gin.Context) {
	data := make([]*models.UserBasic, 10)

	// data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}
