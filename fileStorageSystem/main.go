package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	LoadEnv()
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	h := &Handler{db: db}
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", h.uploadFile)
	r.GET("/getfile", h.getFile)
	r.DELETE("/delete", h.deleteFile)
	r.Run()
}
