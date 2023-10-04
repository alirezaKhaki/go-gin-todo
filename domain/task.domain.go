package domain

import (
	"github.com/alirezaKhaki/go-gin/dto"
	models "github.com/alirezaKhaki/go-gin/model"
)

type ITaskService interface {
	// GetTask(taskId string) (models.Task, error)
	GetTasks(userId int) ([]models.Task, error)
	// UpdateTasks(id string, task models.Task) (models.Task, error)
	// DeleteTask(id string) error
	CreateTask(body dto.CreateTaskRequestBodyDto, userId int) (*models.Task, error)
}
