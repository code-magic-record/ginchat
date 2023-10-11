package models

import (
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

// 自动迁移
// utils.DB.AutoMigrate(&UserBasic{})

func CreateUser(user UserBasic) error {
	result := utils.DB.Create(&user)
	if result.Error != nil {
		// 在这里处理错误，例如输出日志、返回错误信息等
		return result.Error
	}
	return nil
}
