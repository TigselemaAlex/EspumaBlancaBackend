package initializers

import (
	"EspumaBlancaBackend/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func SyncDatabase() {
	migrations()
	createRolesAndAdminUser()
}

func migrations() {
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.User{})
}

func createRolesAndAdminUser() {
	var count int64
	if DB.Model(&models.Role{}).Count(&count); count == 0 {
		createRoles()
	}
	if DB.Model(&models.User{}).Count(&count); count == 0 {
		createAdminUser()
	}
}

func createRoles() {
	fmt.Println("Creating roles...")
	DB.Create(&models.Role{Name: "admin"})
	DB.Create(&models.Role{Name: "user"})
}

func createAdminUser() {
	fmt.Print("Creating admin user...")
	password, err := bcrypt.GenerateFromPassword([]byte("admin12356"), 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	var roles []models.Role
	DB.Find(&roles)
	admin := models.User{Email: "admin@admin.com", Password: string(password), Roles: roles}
	DB.Create(&admin)
}
