package main

import (
	_ "Android_ios/docs"
	"Android_ios/models"
	"Android_ios/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title 牧牛流马(Herding_cattle_and_straying_horses)
// @version 1.0
// @description 用于家畜交易
// @termsOfService http://swagger.io/terms/
// @contact.name phone_number 15294440097
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 8.130.86.26:13000
// @BasePath

func main() {
	// 初始化数据库等
	models.Init()

	// 获取 Gin 路由
	r := router.Router()

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:    ":13000", // 监听的地址和端口
		Handler: r,        // 使用 Gin 路由处理请求
	}

	// 启动 HTTP 服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// 设置超时时间，等待处理已有请求完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	log.Println("Server exiting")
}
