package repository

import (
	"github.com/alirezaKhaki/go-gin/lib"
)

type IUserRepository interface {
	// Create(user *models.User) error
	// FindOne(phoneNumber string) (models.User, error)
}

// UserRepository database structure
type UserRepository struct {
	lib.Database
	logger lib.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db lib.Database, logger lib.Logger) UserRepository {
	return UserRepository{
		Database: db,
		logger:   logger,
	}
}
