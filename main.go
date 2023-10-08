package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	r := router.Router()

	r.Run() // listen and serve on 0.0.0.0:8080
}
