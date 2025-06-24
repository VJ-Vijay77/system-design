package common

import (
	"fmt"
	"socialmedia/database"
	"socialmedia/models"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("/home/vijay/go/src/SystemDesigning/.env")
	if err != nil {
		fmt.Println("No .env file found")
	}
}

func GetServiceContext() *models.Service {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	return &models.Service{Db: db}
}
