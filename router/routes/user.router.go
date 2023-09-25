package router

import (
	"github.com/alirezaKhaki/go-gin/controller"
	"github.com/alirezaKhaki/go-gin/middleware"
	"github.com/alirezaKhaki/go-gin/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Set up routes for the user controller
func SetupUserRoutes(router *gin.RouterGroup, db *gorm.DB) {

	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/", middleware.AuthMiddleware(), userController.FindOne)

		userRoutes.POST("/", userController.Create)
	}
}
