package initializers

import (
	"EspumaBlancaBackend/config"
	"EspumaBlancaBackend/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

func SyncDatabase() {
	db = config.DatabaseConnection()
	migrations()
	createRolesAndAdminUser()
}

func migrations() {
	_ = db.AutoMigrate(&models.Role{})
	_ = db.AutoMigrate(&models.User{})
}

func createRolesAndAdminUser() {
	var db = config.DatabaseConnection()
	var count int64
	if db.Model(&models.Role{}).Count(&count); count == 0 {
		createRoles()
	}
	if db.Model(&models.User{}).Count(&count); count == 0 {
		createAdminUser()
	}
}

func createRoles() {
	fmt.Println("Creating roles...")
	db.Create(&models.Role{Name: "admin"})
	db.Create(&models.Role{Name: "user"})
}

func createAdminUser() {
	fmt.Print("Creating admin user...")
	password, err := bcrypt.GenerateFromPassword([]byte("admin12356"), 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	var roles []models.Role
	db.Find(&roles)
	admin := models.User{Email: "admin@admin.com", Password: string(password), Roles: roles}
	db.Create(&admin)
}
