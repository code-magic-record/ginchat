package main

import (
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	r := router.Router()
	// 自动表迁移
	utils.DB.AutoMigrate(&models.UserBasic{}, &models.UserInfo{})
	r.Run() // listen and serve on 0.0.0.0:8080
}
