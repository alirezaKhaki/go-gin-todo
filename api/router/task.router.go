package router

import (
	"fmt"

	"github.com/alirezaKhaki/go-gin/api/middleware"
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/gin-gonic/gin"
)

type TaskRoutes struct {
	logger  lib.Logger
	handler lib.RequestHandler
	// userController controller.IUserController
	authMiddleware middleware.JWTAuthMiddleware
}

func NewTaskRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	// userController controller.IUserController,
	authMiddleware middleware.JWTAuthMiddleware,
) TaskRoutes {
	return TaskRoutes{
		handler: handler,
		logger:  logger,
		// userController: userController,
		authMiddleware: authMiddleware,
	}
}

func (t TaskRoutes) Setup() {
	taskRoutes := t.handler.ApiGroup.Group("/task")
	{
		taskRoutes.GET("/all", func(ctx *gin.Context) {
			fmt.Println("task")
		})
	}
}
