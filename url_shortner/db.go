package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Define the connection string
	dsn := "host=localhost user=urlshortner password=urlshortner dbname=urlshortner port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// Open a new database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Return the database connection
	return db, nil
}

func (h *dbHandler) SaveToDb(data *Url) (*Url, error) {
	var existingUrl Url
	if err := h.Db.Where("long = ?", data.Long).First(&existingUrl).Error; err == nil {
		return &existingUrl, nil
	}
	if err := h.Db.Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
