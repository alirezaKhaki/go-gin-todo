package domain

import (
	models "github.com/alirezaKhaki/go-gin/model"
	"gorm.io/gorm"
)

type IUserService interface {
	WithTrx(trxHandle *gorm.DB) IUserService
	GetOneUser(id uint) (models.User, error)
	GetAllUser() ([]models.User, error)
	CreateUser(models.User) error
	UpdateUser(models.User) error
	DeleteUser(id uint) error
}
