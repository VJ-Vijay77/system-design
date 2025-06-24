package controllers

import (
	"socialmedia/models"

	"github.com/gin-gonic/gin"
)

type Service struct{
	Config *models.Service
}

func (s *Service) SignUp(c *gin.Context) {
c.JSON(200,"Hello ")
}

func (s *Service) Login(c *gin.Context) {

}
