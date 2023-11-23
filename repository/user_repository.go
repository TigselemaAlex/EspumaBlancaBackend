package repository

import "EspumaBlancaBackend/models"

type UserRepository interface {
	Save(user models.User)
	Update(user models.User)
	Delete(id uint)
	FindById(id uint) (models.User, error)
	FindAll() []models.User
	FindByEmail(email string) (models.User, error)
}
