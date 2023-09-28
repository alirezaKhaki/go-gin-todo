package service

import (
	"github.com/alirezaKhaki/go-gin/domain"
	"github.com/alirezaKhaki/go-gin/dto"
	"github.com/alirezaKhaki/go-gin/lib"
	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/alirezaKhaki/go-gin/repository"
)

// UserService service layer
type UserService struct {
	logger     lib.Logger
	repository repository.UserRepository
	jwtService domain.AuthService
}

// NewUserService creates a new userservice
func NewUserService(logger lib.Logger, repository repository.UserRepository, jwtService domain.AuthService) domain.IUserService {
	return UserService{
		logger:     logger,
		repository: repository,
		jwtService: jwtService,
	}
}

// GetOneUser gets one user
func (s UserService) GetOneUser(id uint) (user models.User, err error) {
	return user, s.repository.Find(&user, id).Error
}

// GetAllUser get all the user
func (s UserService) GetAllUser() (users []models.User, err error) {
	return users, s.repository.Find(&users).Error
}

// CreateUser call to create the user
func (s UserService) CreateUser(body dto.CreateUserRequestBodyDto) (*string, error) {
	user := models.User{Name: body.Name, Email: body.Email, Password: body.Password}
	if err := s.repository.Create(&user).Error; err != nil {
		return nil, err
	}

	token, err := s.jwtService.CreateToken(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// UpdateUser updates the user
func (s UserService) UpdateUser(userId uint, body dto.UpdateUserBodyDto) (*models.User, error) {
	user, err := s.GetOneUser(userId)
	if err != nil {
		return nil, err
	}
	user.Name = body.Name
	err = s.repository.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(id uint) error {
	return s.repository.Delete(&models.User{}, id).Error
}
