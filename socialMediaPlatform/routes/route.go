package routes

import (
	"socialmedia/controllers"
	"socialmedia/models"

	"github.com/gin-gonic/gin"
)


func InitRoutes(r *gin.Engine, s *models.Service) {
	h := &controllers.Service{Config: s}

	r.POST("/signup", h.SignUp)
	r.POST("/login", h.Login)
	r.GET("/user/:id")
	r.GET("/feed")
	r.GET("/like/:id")
	r.POST("/comment")
	r.GET("/follow/:id")
	r.GET("/unfollow/:id")
	r.POST("/post")
	r.DELETE("/post")
}
