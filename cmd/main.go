package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nianfouyi/video-user-api/internal/api"
	"github.com/nianfouyi/video-user-api/internal/models"
	"github.com/nianfouyi/video-user-api/pkg/database"
)

func main() {
	// 初始化数据库连接
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 运行数据库迁移
	err = db.AutoMigrate(&models.User{}, &models.Video{}, &models.Tag{})
	if err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	// 创建 Gin 路由
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// 设置路由
	api.SetupRoutes(router, db)

	// 启动服务器
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
