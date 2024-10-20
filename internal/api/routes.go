package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nianfouyi/video-user-api/internal/api/handlers"
	"github.com/nianfouyi/video-user-api/internal/api/middlewares"
	"github.com/nianfouyi/video-user-api/internal/repositories"
	"github.com/nianfouyi/video-user-api/internal/services"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// 创建仓库
	userRepo := repositories.NewUserRepository(db)
	videoRepo := repositories.NewVideoRepository(db)
	tagRepo := repositories.NewTagRepository(db)
	playbackRecordRepo := repositories.NewPlaybackRecordRepository(db)

	// 创建服务
	userService := services.NewUserService(userRepo)
	videoService := services.NewVideoService(videoRepo)
	tagService := services.NewTagService(tagRepo)
	playbackRecordService := services.NewPlaybackRecordService(playbackRecordRepo)

	// 创建处理器
	userHandler := handlers.NewUserHandler(userService)
	videoHandler := handlers.NewVideoHandler(videoService)
	tagHandler := handlers.NewTagHandler(tagService)
	playbackRecordHandler := handlers.NewPlaybackRecordHandler(playbackRecordService)

	// 公开路由
	public := r.Group("/api")
	{
		public.POST("/users/register", userHandler.Register)
		public.POST("/users/login", userHandler.Login)
		public.POST("/reset-password", userHandler.ResetPassword)
	}

	// 受保护的路由
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		// 用户路由
		protected.PUT("/users/:id", userHandler.UpdateUser)
		protected.DELETE("/users/:id", userHandler.DeleteUser)
		protected.POST("/users/change-password", userHandler.ChangePassword)

		// 视频路由
		protected.POST("/videos", videoHandler.CreateVideos)
		protected.GET("/videos", videoHandler.GetAllVideos)
		protected.GET("/videos/:id", videoHandler.GetVideo)
		protected.PUT("/videos/:id", videoHandler.UpdateVideo)
		protected.DELETE("/videos/:id", videoHandler.DeleteVideo)
		// 播放记录路由
		protected.DELETE("/playback-records", playbackRecordHandler.RecordPlayback)

		protected.POST("/", tagHandler.CreateTag)
		protected.GET("/", tagHandler.GetAllTags)
		protected.GET("/:id", tagHandler.GetTag)
		protected.PUT("/:id", tagHandler.UpdateTag)
		protected.DELETE("/:id", tagHandler.DeleteTag)
	}
}
