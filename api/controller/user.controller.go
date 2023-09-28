package controller

import (
	"net/http"
	"strconv"

	// "github.com/alirezaKhaki/go-gin/constants"
	"github.com/alirezaKhaki/go-gin/domain"
	"github.com/alirezaKhaki/go-gin/dto"
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

// type I

type UserController struct {
	service domain.IUserService
	logger  lib.Logger
}

func NewUserController(userService domain.IUserService, logger lib.Logger) UserController {
	return UserController{
		service: userService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	user, err := u.service.GetOneUser(uint(id))

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// GetUser gets the user
func (u UserController) GetUser(c *gin.Context) {
	users, err := u.service.GetAllUser()
	if err != nil {
		u.logger.Error(err)
	}
	c.JSON(200, gin.H{"data": users})
}

// SaveUser saves the user
func (u UserController) SaveUser(c *gin.Context) {
	var requestBody dto.CreateUserRequestBodyDto

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := u.service.CreateUser(requestBody)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{"tokne": *token})
}

// UpdateUser updates user
func (u UserController) UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{"data": "user updated"})
}

// DeleteUser deletes user
func (u UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := u.service.DeleteUser(uint(id)); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}
