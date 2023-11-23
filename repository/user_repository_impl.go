package repository

import (
	"EspumaBlancaBackend/models"
	"errors"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: Db,
	}
}

func (u UserRepositoryImpl) Save(user models.User) {
	result := u.Db.Create(&user)
	if result.Error != nil {
		panic("error to create the user")
	}
}

func (u UserRepositoryImpl) Update(user models.User) {
	var result = u.Db.Model(&user).Updates(user)
	if result.Error != nil {
		panic("error to update the user")
	}
}

func (u UserRepositoryImpl) Delete(id uint) {
	result := u.Db.Delete(&models.User{}, id)
	if result.Error != nil {
		panic("error to delete the user")
	}
}

func (u UserRepositoryImpl) FindById(id uint) (models.User, error) {
	var user models.User
	result := u.Db.Find(&user, id)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (u UserRepositoryImpl) FindAll() []models.User {
	var users []models.User
	result := u.Db.Find(&users)
	if result.Error != nil {
		panic("error finding users")
	}
	return users
}

func (u UserRepositoryImpl) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := u.Db.Where("email = ?", email).First(&user)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}
