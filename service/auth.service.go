package service

import (
	"errors"

	"github.com/alirezaKhaki/go-gin/domain"
	"github.com/alirezaKhaki/go-gin/lib"
	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/dgrijalva/jwt-go"
)

type JWTAuthService struct {
	env    lib.Env
	logger lib.Logger
}

// NewJWTAuthService creates a new auth service
func NewJWTAuthService(env lib.Env, logger lib.Logger) domain.AuthService {
	return JWTAuthService{
		env:    env,
		logger: logger,
	}
}

// Authorize authorizes the generated token
func (s JWTAuthService) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})
	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("token expired")
		}
	}
	return false, errors.New("couldn't handle token")
}

// CreateToken creates jwt auth tolib
func (s JWTAuthService) CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          user.ID,
		"phoneNumber": user.PhoneNumber,
	})

	tokenString, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		s.logger.Error("JWT validation failed: ", err)
	}

	return tokenString
}
