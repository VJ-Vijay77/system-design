package main

import "github.com/gin-gonic/gin"

type Urlshortner struct {
	Url string `json:"url"`
}

func (h dbHandler) homeHandler(c *gin.Context) {
	c.String(200, "Welcome to the URL Shortener Service!")
}
func (h dbHandler) homeHandler1(c *gin.Context) {
	c.String(200, "Welcome to the URL Shortener Service! Please use the /url endpoint to shorten your URLs.")
}

func (h dbHandler) inputURLHandler(c *gin.Context) {
	var urlShortener Urlshortner
	if err := c.ShouldBindBodyWithJSON(&urlShortener); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	shortUrl := "shrt" + generateRandomString(4)
	dbData := &Url{
		Long:  urlShortener.Url,
		Short: shortUrl,		
	}

	data,err := h.SaveToDb(dbData)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save URL to database","details": err})
		return
	}
	c.JSON(200, gin.H{
		"message": "URL shortened successfully",
		"long": data.Long,
		"short": data.Short,})

}

func (h dbHandler) redirectHandler(c *gin.Context) {
	shortUrl := c.Param("short")
	var url Url
	if err := h.Db.Where("short = ?", shortUrl).First(&url).Error; err != nil {
		c.JSON(404, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(302, url.Long)
}

