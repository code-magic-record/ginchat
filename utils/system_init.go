package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var RDB *redis.Client

func InitRedis(viper *viper.Viper) {
	redisConfig := viper.GetStringMap("redis")
	RDB = redis.NewClient(&redis.Options{
		Addr:     redisConfig["addr"].(string),
		Password: redisConfig["password"].(string),
		DB:       int(redisConfig["db"].(int)),
	})

	_, err := RDB.Ping(RDB.Context()).Result()

	if err != nil {
		fmt.Println(err, "failed to connect redis")
		return
	}
	RDB.Set(RDB.Context(), "ginchat", "test", time.Second)

	fmt.Println("connect redis success")
}

func InitConfig() {
	// 读取配置文件
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	InitMySQL(viper.GetViper())
	InitRedis(viper.GetViper())
}

func InitMySQL(viper *viper.Viper) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{
		Logger: newLogger,
	})

	DB = db
	if err != nil {
		fmt.Println(err, "failed to connect database")
		return
	}
	fmt.Println("connect database success")
}
