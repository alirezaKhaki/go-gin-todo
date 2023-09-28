package domain

import (
	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	Authorize(tokenString string) (*jwt.Token, error)
	CreateToken(models.User) (*string, error)
}
