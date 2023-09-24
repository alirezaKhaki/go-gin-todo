package router

import (
	router "github.com/alirezaKhaki/go-gin/router/routes"
	"github.com/gin-gonic/gin"
	ginDump "github.com/tpkeeper/gin-dump"
)

func NewRouter() {
	server := gin.New()
	server.Use(ginDump.Dump())
	// Create a router group with a common prefix
	apiGroup := server.Group("/api/v1")
	initialRouters(apiGroup)
	server.Run(":5000")
}

func initialRouters(apiGroup *gin.RouterGroup) {
	// Set up routes for the video controller
	router.SetupVideoRoutes(apiGroup)
}
