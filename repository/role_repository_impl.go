package repository

import (
	"EspumaBlancaBackend/models"
	"errors"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewRoleRepositoryImpl(Db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{Db: Db}
}

func (r *RoleRepositoryImpl) Save(role models.Role) {
	result := r.Db.Create(&role)
	if result.Error != nil {
		panic("Error creating role")
	}
}

func (r *RoleRepositoryImpl) FindAll() []models.Role {
	var roles []models.Role
	result := r.Db.Find(&roles)
	if result.Error != nil {
		panic("Error finding roles")
	}
	return roles
}

func (r *RoleRepositoryImpl) FindById(id uint) (models.Role, error) {
	var role models.Role
	result := r.Db.Find(&role, id)
	if result != nil {
		return role, nil
	} else {
		return role, errors.New("role not found")
	}
}
