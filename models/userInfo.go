package models

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"type:varchar(20);not null;unique" json:"username"`
	Address  string `gorm:"type:varchar(20);not null" json:"address"`
}
