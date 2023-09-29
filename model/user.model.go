// models/user.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	//RELATIONS
	Tasks []Task
}

// TableName gives table name of model
func (u User) TableName() string {
	return "users"
}
