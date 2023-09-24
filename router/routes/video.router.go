package router

import (
	"github.com/alirezaKhaki/go-gin/controller"
	"github.com/alirezaKhaki/go-gin/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Set up routes for the user controller
func SetupVideoRoutes(router *gin.RouterGroup, db *gorm.DB) {

	videoService := service.NewVideoService(db)
	videoController := controller.NewVideoController(videoService)

	videoRoutes := router.Group("/video")
	{
		videoRoutes.GET("/all", func(ctx *gin.Context) {
			videos := videoController.FindAll(ctx)
			ctx.JSON(200, videos)
		})

		videoRoutes.POST("/", func(ctx *gin.Context) {
			videos := videoController.Save(ctx)
			ctx.JSON(200, videos)
		})
	}
}
