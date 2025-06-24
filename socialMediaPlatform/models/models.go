package models

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Password string
	Email    string `gorm:"unique"`
}

type Post struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	Name      string
	Caption   string
	Like      int
	Comment   int
}

type Like struct {
	UserID uint
	PostID uint
}

type Comment struct{
	UserID uint
	PostID uint
	Comment string
}

type Followers struct {
	UserID uint
	Followers int
	Following int
}