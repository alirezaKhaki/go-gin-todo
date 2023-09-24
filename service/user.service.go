package service

import (
	"fmt"

	dto "github.com/alirezaKhaki/go-gin/dto"
	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/jinzhu/gorm"
)

type IUserService interface {
	Create(user dto.CreateUserRequestBodyDto) (string, error)
	FindOne(phoneNumber string) (models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) IUserService {
	return &userService{db: db}
}

func (service *userService) Create(body dto.CreateUserRequestBodyDto) (string, error) {
	err := service.db.Create(&models.User{Name: body.Name, PhoneNumber: body.PhoneNumber, Password: body.Password}).Error
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	return "token", nil
}

func (service *userService) FindOne(phoneNumber string) (models.User, error) {
	var user models.User
	err := service.db.Where(`"phoneNumber" = ?`, phoneNumber).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}
