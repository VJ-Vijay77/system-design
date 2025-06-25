package database

import (
	"os"
	"socialMediaPlatform/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DSN_SOCIAL")
	// Open a new database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Comment{}, &models.Followers{}, &models.Like{}, &models.Post{}, &models.User{})

	// Return the database connection
	return db
}
