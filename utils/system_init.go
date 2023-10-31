package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/natefinch/lumberjack"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var RDB *redis.Client

func InitSystemConfig() {
	viper := getYmlConfig()
	InitMySQL(viper)
	InitRedis(viper)
}

func getYmlConfig() *viper.Viper {
	// 读取配置文件
	viper.SetConfigName("app")
	// viper.AddConfigPath("config")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config" // 使用默认值
	}
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	return viper.GetViper()
}

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
	go func() {
		ticker := time.NewTicker(time.Second * 60 * 10)
		for range ticker.C {
			log.Default().Println("定时查询redis，保证redis连接不断开")
			RDB.Set(RDB.Context(), "ginchat", "test", time.Second)
		}
	}()
	fmt.Println("connect redis success")
}

func InitMySQL(viper *viper.Viper) {

	// 创建一个日志文件
	file := &lumberjack.Logger{
		Filename:   "logs/gorm.log", // 日志文件路径
		MaxSize:    1,               // 单个日志文件最大容量（以 MB 为单位）
		MaxBackups: 30,              // 保留旧文件的最大数量
		MaxAge:     7,               // 旧文件保留的最大天数
		Compress:   true,            // 是否启用压缩
	}

	newLogger := logger.New(
		// log.New(file, "\r\n", log.LstdFlags), // 将日志写入文件
		log.New(file, "\r\n", log.LstdFlags), // io writer
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
	token := CreateToken(map[string]interface{}{
		"test": "test",
	})

	go func() {
		ticker := time.NewTicker(time.Second * 60 * 10)
		for range ticker.C {
			fmt.Println("定时查询数据库，保证数据库连接不断开")
			// 在这里添加需要轮询的逻辑
			DB.Exec("select 1")
		}
	}()

	fmt.Println(token, "token")
}
