package domain

import (
	"github.com/alirezaKhaki/go-gin/dto"
	models "github.com/alirezaKhaki/go-gin/model"
)

type IUserService interface {
	GetOneUser(id uint) (models.User, error)
	GetAllUser() ([]models.User, error)
	CreateUser(dto.CreateUserRequestBodyDto) (*string, error)
	UpdateUser(userId uint, body dto.UpdateUserBodyDto) (*models.User, error)
	DeleteUser(id uint) error
}
