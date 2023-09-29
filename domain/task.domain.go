package domain

import models "github.com/alirezaKhaki/go-gin/model"

type ITaskService interface {
	GetTask(taskId string) (models.Task, error)
	GetTasks() ([]models.Task, error)
	UpdateTasks(id string, task models.Task) (models.Task, error)
	DeleteTask(id string) error
	CreateTask(task models.Task) (string, error)
}
