package controller

import (
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	GetAllTasks(c *gin.Context)
}

type TaskController struct {
	logger lib.Logger
}

func NewTaskController(logger lib.Logger) ITaskController {
	return TaskController{
		logger: logger,
	}
}

func (t TaskController) GetAllTasks(c *gin.Context) {

	c.JSON(200, gin.H{"data": "hello"})
}
