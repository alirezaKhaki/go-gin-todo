// models/user.go
package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string
	PhoneNumber string `gorm:"column:phoneNumber"`
	Password    string
}

// UserClaims represents the claims in a JWT token.
type UserClaims struct {
	UserID      uint   `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	jwt.StandardClaims
}
