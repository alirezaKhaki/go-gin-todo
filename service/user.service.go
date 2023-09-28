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
	jwtService JWTAuthService
}

// NewUserService creates a new userservice
func NewUserService(logger lib.Logger, repository repository.UserRepository) domain.IUserService {
	return UserService{
		logger:     logger,
		repository: repository,
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
func (s UserService) CreateUser(dto.CreateUserRequestBodyDto) (*string, error) {
	var user models.User
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
func (s UserService) UpdateUser(user models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(id uint) error {
	return s.repository.Delete(&models.User{}, id).Error
}
