package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (h *Handler) uploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")

	path := os.Getenv("FILESTORAGE")
	fmt.Println("path ====== ", path)
	savePath := filepath.Join(path, file.Filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(200, gin.H{"msg": "Error uploading file", "err": err})
		return
	}
	fileForDb := &File{
		FileName: file.Filename,
		FilePath: path,
		FileSize: file.Size,
	}
	h.saveFileToDb(fileForDb)
	c.String(200, "file uploaded successfully: "+file.Filename)
}

func (h *Handler) getFile(c *gin.Context) {
	c.JSON(200, "ok")
}

func (h *Handler) deleteFile(c *gin.Context) {
	var filename FileName
	var file File
	fmt.Println("Method:::::", c.Request.Method)


	c.ShouldBindJSON(&filename)
	if err := h.db.Where("file_name=?", filename.FileName).First(&file).Error; err != nil {
		fmt.Println("db error", err)
		c.JSON(400, err)
		return
	}

		if err := os.Remove(file.FilePath+file.FileName); err != nil {
			if os.IsNotExist(err) {
				c.JSON(http.StatusNotFound, "file not found")
				return
			} else {
				c.JSON(500, err)
			}
			return
		}
	

	if err := h.deleteFileFromDb(filename.FileName); err != nil {
		c.JSON(400, gin.H{
			"msg": "file deletion failed",
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{"msg":"file deletion successful"})
}
