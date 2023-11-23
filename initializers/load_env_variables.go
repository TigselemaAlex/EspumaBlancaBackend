package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	fmt.Println("ENV")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
