package main

import (
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitSystemConfig()
	r := router.Router()
	// 自动表迁移
	utils.DB.AutoMigrate(&models.UserBasic{})
	r.Run("localhost:8888") // listen and serve on 0.0.0.0:8888
}
