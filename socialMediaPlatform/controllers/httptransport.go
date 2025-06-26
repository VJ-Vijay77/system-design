package controllers

import (
	"errors"
	"net/http"
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
		c.JSON(400, gin.H{
			"msg":   "Error in binding",
			"error": err.Error(),
		})
		return
	}
	if len(user.Password) < 8 {
		c.JSON(400, "Password must have atleast 8 characters")
		return
	}
	hashedPassword, err := s.Common.HashPassword(user.Password)
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
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "User Successfully Created, Goto Login")
}

func (s *Service) Login(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"msg":   "Error in binding",
			"error": err.Error(),
		})
		return
	}
	userFromDb := s.getPassword(user.UserID)
	pass := s.Common.GetPasswordFromHash(userFromDb.Password, user.Password)
	if !pass {
		c.JSON(http.StatusUnauthorized, "Wrong username or password")
		return
	}

	token, err := s.Common.GenerateJwtToken(userFromDb.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "failed to generate jwt token",
			"err": err.Error(),
		})
	}

	c.JSON(200, token)
}

func (s *Service) GetUserID(c *gin.Context) {
	userID := c.Param("id")
	var user models.User
	if err := s.Db.Where("name=?", userID).Or("email=?", userID).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, "User does not exist")
			return
		}
		c.JSON(500, err)
		return
	}
	c.JSON(200, gin.H{
		"Name":  user.Name,
		"Email": user.Email,
	})
}
