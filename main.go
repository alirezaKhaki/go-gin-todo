package main

import (
	"github.com/alirezaKhaki/go-gin/database"
	"github.com/alirezaKhaki/go-gin/router"
)

func main() {
	// Initialize the database
	_, err := database.Init()
	if err != nil {
		panic(err)
	}
	defer database.DB.Close()

	router.NewRouter(database.DB)
}
