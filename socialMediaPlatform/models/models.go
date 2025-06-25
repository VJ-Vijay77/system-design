package models

import (
	"time"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Password string
	Email    string `gorm:"unique"`
}

type UserSignup struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
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

type Comment struct {
	UserID  uint
	PostID  uint
	Comment string
}

type Followers struct {
	UserID    uint
	Followers int
	Following int
}
