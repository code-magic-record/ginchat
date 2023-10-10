package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	// fmt.Println("config app:", viper.Get("mysql.dns"))
	InitMySQL(viper.GetViper())
}

func InitMySQL(viper *viper.Viper) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // 日志级别
			Colorful:      true,          // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{
		Logger: newLogger,
	})

	DB = db
	if err != nil {
		fmt.Println(err, "failed to connect database")
	}
	fmt.Println("connect database success")

}
