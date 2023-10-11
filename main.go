package main

import (
	"ginchat/models"
	"ginchat/router"
	"ginchat/service"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	r := router.Router()
	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET("/index", service.GetIndex)
		}
	}
	// 自动表迁移
	utils.DB.AutoMigrate(&models.UserBasic{}, &models.UserInfo{})
	r.Run() // listen and serve on 0.0.0.0:8080
}
