package service

import (
	"EspumaBlancaBackend/data/request"
	"EspumaBlancaBackend/data/response"
)

type UserService interface {
	SignIn(loginRequest request.LoginRequest) (response.UserResponse, error)
	Create(user request.UserRequest)
	Update(user request.UserRequest)
	Delete(id uint)
	FindById(id uint) response.UserResponse
	FindAll() []response.UserResponse
}
