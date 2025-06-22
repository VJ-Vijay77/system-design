package main



type File struct {
	ID uint `gorm:"PrimaryKey"`
	FileName string `gorm:"uniqueKey"`
	FilePath string
	FileSize int64
}

type FileName struct{
	FileName string `json:"filename"`
}