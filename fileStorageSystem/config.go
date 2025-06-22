package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("/home/vijay/go/src/SystemDesigning/.env")
	if err != nil {
		fmt.Println("No .env file found")
	}
}
