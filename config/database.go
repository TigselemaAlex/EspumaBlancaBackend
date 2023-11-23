package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DatabaseConnection() *gorm.DB {
	stringConnection := os.Getenv("DB_STRING_CONNECTION")
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	return db
}
