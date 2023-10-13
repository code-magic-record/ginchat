package main

import (
	"ginchat/models"
	"ginchat/utils"
)

func main() {
	// db, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// },
	utils.InitSystemConfig()

	// Migrate the schema
	utils.DB.AutoMigrate(&models.UserBasic{})

	// // Create
	user := &models.UserBasic{}
	user.Name = "zyg"

	utils.DB.Create(user)

	// // Read
	// fmt.Println(utils.DB.First(user, 1)) // find product with integer primary key

	// update
	// db.Model(user).Update("Name", "zyg2")
}
