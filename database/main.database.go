package database

import (
	"context"

	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/fx"
)

var DB *gorm.DB

type Database struct{}

func NewDatabase(lc fx.Lifecycle) *gorm.DB {
	// Initialize the database
	_, err := Init()
	if err != nil {
		panic(err)
	}
	// Register a function to close the database when the application shuts down
	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return DB.Close()
		},
	})

	return DB
}

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
