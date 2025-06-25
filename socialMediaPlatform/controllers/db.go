package controllers

import "socialMediaPlatform/models"

func (h *Service) saveUser(params *models.User) error {
	if err := h.Db.Create(&params).Error; err != nil {
		return err
	}
	return nil
}
