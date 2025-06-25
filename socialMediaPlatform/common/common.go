package common

import (
	"fmt"

	"github.com/ansel1/merry"
	"github.com/jackc/pgconn"
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

func (s *Service) GetPasswordFromHash(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return merry.New("wrong password")
	}
	return nil
}

func (s *Service) IsUniqueConstraintError(err error) bool {
	// PostgreSQL
	if pgErr, ok := err.(*pgconn.PgError); ok {
		return pgErr.Code == "23505"
	}
	return false
}