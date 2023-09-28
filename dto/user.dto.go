package dto

type CreateUserRequestBodyDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type FindOneUserRequestBodyDto struct {
	PhoneNumber string `json:"phoneNumber" binding:"required" form:"phoneNumber"`
}
