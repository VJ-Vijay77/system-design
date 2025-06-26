package common

import (
	"fmt"
	"os"
	"time"

	"github.com/ansel1/merry"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func LoadEnv() {
	err := godotenv.Load("/home/vijay/go/src/SystemDesigning/.env")
	if err != nil {
		fmt.Println("No .env file found")
	}
}

func (s *Service) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", merry.Append(err, "unalbe to get the hash")
	}
	return string(hash), nil
}

func (s *Service) GetPasswordFromHash(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

func (s *Service) GenerateJwtToken(userid uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userid,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
