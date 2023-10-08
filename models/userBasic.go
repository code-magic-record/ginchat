package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name         string
	Password     string
	Phone        string
	Email        string
	Identity     string
	ClinetIP     string
	ClintProt    string
	LoginTime    string
	Heartbeat    string // 考虑用户心跳时间
	LoginOutTime string `gorm:"column:login_out_time" json: "login_out_time"`
	IsLogout     bool   // 考虑用户是否退出
	DeviceInfo   string // 考虑用户设备信息
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}
