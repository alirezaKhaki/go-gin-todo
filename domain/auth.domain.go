package domain

import models "github.com/alirezaKhaki/go-gin/model"

type AuthService interface {
	Authorize(tokenString string) (bool, error)
	CreateToken(models.User) (*string, error)
}
