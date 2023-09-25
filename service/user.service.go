package service

import (
	"errors"

	dto "github.com/alirezaKhaki/go-gin/dto"
	models "github.com/alirezaKhaki/go-gin/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(user dto.CreateUserRequestBodyDto) (string, error)
	FindOne(phoneNumber string) (models.User, error)
	GenerateToken(user models.User) (string, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) IUserService {
	return &userService{db: db}
}

func (service *userService) Create(body dto.CreateUserRequestBodyDto) (string, error) {
	user, err := service.FindOne(body.PhoneNumber)
	if err != nil {
		return "error", err
	}

	if len(user.PhoneNumber) > 0 {
		return "", errors.New("user already exist")
	}

	password := []byte(body.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	newUser := &models.User{
		Name:        body.Name,
		PhoneNumber: body.PhoneNumber,
		Password:    string(hashedPassword),
	}

	err = service.db.Create(newUser).Error
	if err != nil {
		return "", err
	}

	createdUser := *newUser

	token, err := service.GenerateToken(createdUser)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *userService) FindOne(phoneNumber string) (models.User, error) {
	var user models.User
	if err := service.db.Where(`"phoneNumber" = ?`, phoneNumber).First(&user).Error; err != nil {
		// Handle the error, user not found, or other issues
		// You can check err for gorm.ErrRecordNotFound to specifically handle "user not found" scenarios
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return user, err
		}
	}

	return user, nil
}

func (service *userService) GenerateToken(user models.User) (string, error) {
	claims := models.UserClaims{UserID: user.ID, PhoneNumber: user.PhoneNumber}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("your-secret-key")

	return token.SignedString(secretKey)
}
