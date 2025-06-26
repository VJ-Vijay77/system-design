package controllers

import (
	"socialMediaPlatform/models"
	"strings"

	"github.com/ansel1/merry"
)

func (s *Service) saveUser(params *models.User) error {
	if err := s.Db.Create(&params).Error; err != nil {
		msg := err.Error()

		if strings.Contains(msg, "duplicate") {
			return merry.New("User already exist")
		}
		return err
	}
	return nil
}

func (s *Service) getPassword(userid string) *models.User{
	var user models.User
	if err := s.Db.Where("name = ?",userid).Or("email=?",userid).Find(&user).Error; err != nil {
		return nil
	}
	return &user
}
