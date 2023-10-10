package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;unique" json:"name"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
}

func (UserBasic) TableName() string {
	return "user_basics"
}

func CreateUser(user UserBasic) error {
	// 自动迁移
	utils.DB.AutoMigrate(&UserBasic{})
	// utils.DB.Create(&user)
	result := utils.DB.Create(&user)
	fmt.Println(result, "创建成功了？")
	if result.Error != nil {
		// 在这里处理错误，例如输出日志、返回错误信息等
		return result.Error
	}
	return nil
}
