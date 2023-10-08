package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	// 读取配置文件
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config app:", viper.Get("mysql.dns"))
}

func InitMySQL() {
	DB, _ := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})

	fmt.Println(DB)
}
