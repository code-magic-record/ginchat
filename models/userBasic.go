package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Id       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(20);not null;unique" json:"name"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Phone    string `gorm:"type:varchar(20);not null;unique" json:"phone"`
	Avatar   string `gorm:"type:varchar(255); null" json:"avatar"`
}

func (UserBasic) TableName() string {
	return "user_basics"
}

func CreateUser(user UserBasic) error {
	result := utils.DB.Create(&user)
	if result.Error != nil {
		// 在这里处理错误，例如输出日志、返回错误信息等
		return result.Error
	}
	return nil
}

func SearchPhone(phone string) bool {
	var user UserBasic
	utils.DB.Where("phone = ?", phone).First(&user)
	return user.Id != 0
}

func SearchUserByPhone(phone string) UserBasic {
	var user UserBasic
	utils.DB.Where("phone = ?", phone).First(&user)
	return user
}
