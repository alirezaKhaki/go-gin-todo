package lib

import (
	"fmt"

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

	return Database{
		DB: db,
	}
}
