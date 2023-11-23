package service

import (
	"EspumaBlancaBackend/data/request"
	"EspumaBlancaBackend/data/response"
	"EspumaBlancaBackend/models"
	"EspumaBlancaBackend/repository"
	"github.com/go-playground/validator/v10"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	Validate       *validator.Validate
}

func NewRoleServiceImpl(roleRepository repository.RoleRepository, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		Validate:       validate,
	}
}

func (r RoleServiceImpl) Create(roleReq request.RoleRequest) {
	err := r.Validate.Struct(roleReq)
	if err != nil {
		panic("error to create the role")
	}
	role := models.Role{
		Name: roleReq.Name,
	}
	r.RoleRepository.Save(role)
}

func (r RoleServiceImpl) FindAll() []response.RoleResponse {
	result := r.RoleRepository.FindAll()
	var roles []response.RoleResponse
	for _, value := range result {
		role := response.RoleResponse{
			Id:   int(value.ID),
			Name: value.Name,
		}
		roles = append(roles, role)
	}
	return roles
}

func (r RoleServiceImpl) FindById(id uint) response.RoleResponse {
	result, err := r.RoleRepository.FindById(id)
	if err != nil {
		panic(err)
	}
	role := response.RoleResponse{
		Name: result.Name,
	}
	return role
}
