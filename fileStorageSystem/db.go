package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")
	// Open a new database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&File{})

	// Return the database connection
	return db, nil
}

func (h *Handler) saveFileToDb(file *File) error {
	if err := h.db.Create(&file).Error; err != nil {
		fmt.Println("db error", err)
		return err
	}
	return nil
}


func (h *Handler) deleteFileFromDb(filename string) error{
	if err := h.db.Where("file_name=?",filename).Delete(File{}).Error; err != nil {
		fmt.Println("db error",err)
		return err
	}
	return nil
}