package service

import (
	"github.com/alirezaKhaki/go-gin/domain"
	"github.com/alirezaKhaki/go-gin/lib"
	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/alirezaKhaki/go-gin/repository"
)

type TaskService struct {
	logger     lib.Logger
	repository repository.ITaskRepository
}

func NewTaskService(logger lib.Logger, repo repository.ITaskRepository) domain.ITaskService {
	return TaskService{
		logger:     logger,
		repository: repo,
	}
}

func (s TaskService) GetTasks(userId int) ([]models.Task, error) {
	tasks, err := s.repository.GetAll(uint(userId))
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// func (s TaskService) CreateTask(task models.Task) (string, error) {

// }
