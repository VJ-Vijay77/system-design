package database

import (
	"fmt"
	"os"
	"socialmedia/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DSN_SOCIAL")
	fmt.Println("Database dsn=====",dsn)
	// Open a new database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Comment{}, &models.Followers{}, &models.Like{}, &models.Post{}, &models.User{})

	// Return the database connection
	return db, nil
}
