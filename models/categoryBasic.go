package models

import "ginchat/utils"

type CategoryBasic struct {
	Id       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(20);not null;unique" json:"name"`
	ImaeUrl  string `gorm:"type:varchar(255); null" json:"image_url"`
	Describe string `gorm:"type:varchar(255); null" json:"describe"`
}

func (CategoryBasic) TableName() string {
	return "category_basics"
}

func CreateCategory(category CategoryBasic) error {
	result := utils.DB.Create(&category)
	if result.Error != nil {
		// 在这里处理错误，例如输出日志、返回错误信息等
		return result.Error
	}
	return nil
}
