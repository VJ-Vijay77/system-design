package routes

import (
	"socialMediaPlatform/di"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, c *di.Config) {

	r.POST("/signup", c.Config.SignUp)
	r.POST("/login", c.Config.Login)

	authGroup := r.Group("/user")
	authGroup.Use(c.Config.AuthMiddleWare())


	authGroup.GET("/:id",c.Config.GetUserID)
	authGroup.GET("/feed")
	authGroup.GET("/like/:id")
	authGroup.POST("/comment")
	authGroup.GET("/follow/:id")
	authGroup.GET("/unfollow/:id")
	authGroup.POST("/post")
	authGroup.DELETE("/post")
}
