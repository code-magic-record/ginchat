package main

import (
	"ginchat/router"
	"ginchat/service"
	"ginchat/utils"
	// docs is generated by Swag CLI, you have to import it.
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
	r.Run() // listen and serve on 0.0.0.0:8080
}
