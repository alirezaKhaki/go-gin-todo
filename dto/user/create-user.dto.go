package dto

type CreateUserRequestBodyDto struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
