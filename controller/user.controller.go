package controller

import (
	"net/http"

	"github.com/alirezaKhaki/go-gin/dto"
	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/alirezaKhaki/go-gin/service"
	"github.com/gin-gonic/gin"
)

type IUserContoller interface {
	Create(ctx *gin.Context)
	FindOne(ctx *gin.Context)
}

type userController struct {
	service service.IUserService
}

// Define a custom result type
type Result struct {
	Value models.User
	Err   error
}

func NewUserController(service service.IUserService) IUserContoller {
	return &userController{
		service: service,
	}
}

func (c *userController) FindOne(ctx *gin.Context) {
	// Retrieve the user from the context
	claims, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}

	// Access the user as needed
	customClaims, ok := claims.(*models.UserClaims)
	if !ok {
		// Handle unexpected claim type
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid claim type"})
		return
	}

	// Access the phone number
	phoneNumber := customClaims.PhoneNumber

	user, err := c.service.FindOne(phoneNumber)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(user.PhoneNumber) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (c *userController) Create(ctx *gin.Context) {
	var requestBody dto.CreateUserRequestBodyDto

	// Bind the request body to the struct
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Create(requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
