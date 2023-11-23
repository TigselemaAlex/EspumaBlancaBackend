package service

import (
	"EspumaBlancaBackend/data/request"
	"EspumaBlancaBackend/data/response"
)

type RoleService interface {
	Create(roleReq request.RoleRequest)
	FindAll() []response.RoleResponse
	FindById(id uint) response.RoleResponse
}
