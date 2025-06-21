package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type dbHandler struct {
	Db *gorm.DB
}

func main() {
	r := gin.Default()
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	db.AutoMigrate(&Url{})

	h := &dbHandler{
		Db: db,
	}


	r.GET("/", h.homeHandler)
	r.GET("/home", h.homeHandler1)
	r.POST("/url", h.inputURLHandler)
	r.GET("/short/:short", h.redirectHandler)
	r.Run()
}
