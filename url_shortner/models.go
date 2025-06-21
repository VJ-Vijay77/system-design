package main

type Url struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Long  string `json:"long" gorm:"unique"`
	Short string `json:"short" gorm:"unique"`
}