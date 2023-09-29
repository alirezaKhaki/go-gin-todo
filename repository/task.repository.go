package repository

import (
	"github.com/alirezaKhaki/go-gin/lib"
	models "github.com/alirezaKhaki/go-gin/model"
)

type ITaskRepository interface {
	GetAll(userId uint) ([]models.Task, error)
}

type TaskRepository struct {
	lib.Database
	lib.Logger
}

func NewTaskRepository(db lib.Database, logger lib.Logger) ITaskRepository {
	return TaskRepository{
		db,
		logger,
	}
}

func (t TaskRepository) GetAll(userId uint) ([]models.Task, error) {
	var tasks []models.Task
	
	if err := t.Model(&models.Task{UserID: (userId)}).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
