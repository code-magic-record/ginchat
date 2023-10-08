package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "zyg"

	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // find product with integer primary key

	// update
	// db.Model(user).Update("Name", "zyg2")
}
