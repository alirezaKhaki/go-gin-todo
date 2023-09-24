package database

import (
	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=go-todo sslmode=disable password=postgres")
	if err != nil {
		return nil, err
	}
	DB = db

	migrate(db)

	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
