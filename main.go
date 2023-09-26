package main

import (
	"github.com/alirezaKhaki/go-gin/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	bootstrap.RunServer(bootstrap.CommonModules)
}
