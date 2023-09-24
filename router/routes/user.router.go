package router

import (
	"net/http"

	"github.com/alirezaKhaki/go-gin/controller"
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
		userRoutes.GET("/findOne", func(ctx *gin.Context) {
			user := userController.FindOne(ctx)
			if user.Err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": user.Err.Error()})

			}
			ctx.JSON(200, user.Value)
		})

		userRoutes.POST("/", func(ctx *gin.Context) {
			token, err := userController.Create(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(200, token)
		})
	}
}
