package models

type Role struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"unique"`
}
