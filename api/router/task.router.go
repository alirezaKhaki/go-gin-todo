package router

import (
	"github.com/alirezaKhaki/go-gin/api/controller"
	"github.com/alirezaKhaki/go-gin/api/middleware"
	"github.com/alirezaKhaki/go-gin/lib"
)

type TaskRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	taskController controller.ITaskController
	authMiddleware middleware.JWTAuthMiddleware
}

func NewTaskRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	taskController controller.ITaskController,
	authMiddleware middleware.JWTAuthMiddleware,
) TaskRoutes {
	return TaskRoutes{
		handler:        handler,
		logger:         logger,
		taskController: taskController,
		authMiddleware: authMiddleware,
	}
}

func (t TaskRoutes) Setup() {
	taskRoutes := t.handler.ApiGroup.Group("/task")
	{
		taskRoutes.GET("/all", t.authMiddleware.Handler(), t.taskController.GetAllTasks)
		taskRoutes.POST("/", t.authMiddleware.Handler(), t.taskController.CreateTask)
	}

}
