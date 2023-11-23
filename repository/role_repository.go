package repository

import "EspumaBlancaBackend/models"

type RoleRepository interface {
	Save(role models.Role)
	FindAll() []models.Role
	FindById(id uint) (models.Role, error)
}
