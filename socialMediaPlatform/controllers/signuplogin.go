package controllers

import (
	"socialMediaPlatform/common"
	"socialMediaPlatform/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	Db     *gorm.DB
	Common *common.Service
}

func (s *Service) SignUp(c *gin.Context) {
	var user models.UserSignup
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(400, "Error in binding json")
		return
	}
	if len(user.Password) < 8 {
		c.JSON(400, "Password must have atleast 8 characters")
		return
	}
	hashedPassword,err := s.Common.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, "Internal Servor error:  Hash issue")
	}
	userToDb := &models.User{
		Name:     user.Name,
		Password: hashedPassword,
		Email:    user.Email,
	}

	err = s.saveUser(userToDb)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "User Successfully Created, goto login")
}

func (s *Service) Login(c *gin.Context) {

}
