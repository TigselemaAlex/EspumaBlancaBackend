package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Roles    []Role `gorm:"many2many:user_roles;not null"`
}
