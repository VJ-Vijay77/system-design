package routes

import (
	"socialMediaPlatform/di"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, c *di.Config) {

	r.POST("/signup", c.Config.SignUp)
	r.POST("/login", c.Config.Login)
	r.GET("/user/:id")
	r.GET("/feed")
	r.GET("/like/:id")
	r.POST("/comment")
	r.GET("/follow/:id")
	r.GET("/unfollow/:id")
	r.POST("/post")
	r.DELETE("/post")
}
