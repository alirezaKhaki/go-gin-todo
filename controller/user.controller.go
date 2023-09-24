package controller

import (
	dto "github.com/alirezaKhaki/go-gin/dto/user"
	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/alirezaKhaki/go-gin/service"
	"github.com/gin-gonic/gin"
)

type IUserContoller interface {
	Create(ctx *gin.Context) (string, error)
	FindOne(ctx *gin.Context) Result
}

type userController struct {
	service service.IUserService
}

// Define a custom result type
type Result struct {
	Value models.User
	Err   error
}

type PhoneNumberRequestBody struct {
	PhoneNumber string `json:"phoneNumber"`
}

func NewUserController(service service.IUserService) IUserContoller {
	return &userController{
		service: service,
	}
}

func (c *userController) FindOne(ctx *gin.Context) Result {
	var requestBody PhoneNumberRequestBody

	// Bind the request body to the struct
	if err := ctx.Bind(&requestBody); err != nil {
		return Result{models.User{}, err}
	}

	// Access the phone number
	phoneNumber := requestBody.PhoneNumber

	user, err := c.service.FindOne(phoneNumber)
	if err != nil {
		return Result{models.User{}, err}
	}
	return Result{Value: user, Err: nil}
}

func (c *userController) Create(ctx *gin.Context) (string, error) {
	var requestBody dto.CreateUserRequestBodyDto

	// Bind the request body to the struct
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		return "", err
	}

	token, err := c.service.Create(requestBody)
	if err != nil {
		return "", err
	}
	return token, nil

}
