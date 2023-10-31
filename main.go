package main

import (
	"context"
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	utils.InitSystemConfig()
	r := router.Router()
	// 自动表迁移
	utils.DB.AutoMigrate(&models.UserBasic{}, &models.CategoryBasic{})
	// r.Run("0.0.0.0:8888") // listen and serve on 0.0.0.0:8888

	srv := &http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
