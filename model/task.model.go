package models

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	DueDate     time.Time `json:"dueDate"`
}
