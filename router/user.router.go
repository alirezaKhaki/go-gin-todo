package router

import (
	"github.com/alirezaKhaki/go-gin/controller"
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/alirezaKhaki/go-gin/middleware"
)

type UserRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController controller.UserController
	authMiddleware middleware.JWTAuthMiddleware
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController controller.UserController,
	authMiddleware middleware.JWTAuthMiddleware,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}

// Set up routes for the user controller
func (u UserRoutes) Setup() {
	userRoutes := u.handler.Gin.Group("/user")
	{
		userRoutes.GET("/", u.authMiddleware.Handler(), u.userController.GetOneUser)

		userRoutes.POST("/", u.userController.SaveUser)
	}
}
