package lib

import (
	"fmt"

	models "github.com/alirezaKhaki/go-gin/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct{ *gorm.DB }

func NewDatabase(env Env, logger Logger) Database {

	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	dbname := env.DBName

	url := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, dbname, password)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	DB = db

	logger.Info("Database connection established")

	migrate(env, logger, DB)

	return Database{
		DB: db,
	}
}

func migrate(env Env, logger Logger, db *gorm.DB) {
	if env.Environment == "development" {
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Task{})
		logger.Info("Database migration done")
	}
}
