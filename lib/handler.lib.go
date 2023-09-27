package lib

import (
	"github.com/gin-gonic/gin"
)

// RequestHandler function
type RequestHandler struct {
	Gin      *gin.Engine
	ApiGroup *gin.RouterGroup
}

// NewRequestHandler creates a new request handler
func NewRequestHandler(logger Logger) RequestHandler {
	gin.DefaultWriter = logger.GetGinLogger()
	engine := gin.New()
	apiGroup := engine.Group("/api/v1")
	return RequestHandler{Gin: engine, ApiGroup: apiGroup}
}
