package controller

import (
	"net/http"

	"github.com/alirezaKhaki/go-gin/domain"
	"github.com/alirezaKhaki/go-gin/dto"
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	GetAllTasks(c *gin.Context)
	CreateTask(c *gin.Context)
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

func (t TaskController) CreateTask(c *gin.Context) {
	var body dto.CreateTaskRequestBodyDto

	userID := c.MustGet("id").(int)

	if userID <= 0 {
		c.JSON(400, gin.H{"error": "invalid token"})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, err := t.service.CreateTask(body, userID)
	if err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": &task,
	})
}
