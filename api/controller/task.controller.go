package controller

import (
	"github.com/alirezaKhaki/go-gin/domain"
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	GetAllTasks(c *gin.Context)
}

type TaskController struct {
	logger  lib.Logger
	service domain.ITaskService
}

func NewTaskController(logger lib.Logger, service domain.ITaskService) ITaskController {
	return TaskController{
		logger:  logger,
		service: service,
	}
}

func (t TaskController) GetAllTasks(c *gin.Context) {
	userID := c.MustGet("id").(int)

	if userID <= 0 {
		c.JSON(400, gin.H{"error": "invalid token"})
		return
	}

	tasks, err := t.service.GetTasks(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": tasks})
}
