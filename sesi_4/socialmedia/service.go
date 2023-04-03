package socialmedia

import (
	"sesi_4_final_project/models"
	"time"
)

type Service interface {
	CreateSocialMedia(input CreateSocialMediaInput, userID uint) (models.SocialMedia, error)
	GetAll(userID uint) ([]models.SocialMedia, error)
	GetOne(socialMediaID uint) (models.SocialMedia, error)
	UpdateSocialMedia(socialMediaID uint, input CreateSocialMediaInput, userID uint) (models.SocialMedia, error)
	DeleteSocialMedia(socialMediaID uint, userID uint) error
	GetOneByUserID(socialMediaID uint, userID uint) (models.SocialMedia, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateSocialMedia(input CreateSocialMediaInput, userID uint) (models.SocialMedia, error) {
	socialmedia := models.SocialMedia{}
	socialmedia.Name = input.Name
	socialmedia.SocialMediaUrl = input.SocialMediaUrl
	socialmedia.UserID = userID

	newSocialMedia, err := s.repository.Save(socialmedia)
	if err != nil {
		return newSocialMedia, err
	}

	return newSocialMedia, nil
}

func (s *service) GetAll(userID uint) ([]models.SocialMedia, error) {
	socialmedias, err := s.repository.FindAll(userID)
	if err != nil {
		return []models.SocialMedia{}, err
	}

	return socialmedias, nil
}

func (s *service) GetOneByUserID(socialMediaID uint, userID uint) (models.SocialMedia, error) {
	socialmedia, err := s.repository.FindByUserID(socialMediaID, userID)
	if err != nil {
		return models.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (s *service) GetOne(socialMediaID uint) (models.SocialMedia, error) {
	socialmedia, err := s.repository.FindByID(socialMediaID)
	if err != nil {
		return models.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (s *service) UpdateSocialMedia(socialMediaID uint, inputData CreateSocialMediaInput, userID uint) (models.SocialMedia, error) {
	socialmedia, err := s.repository.FindByUserID(socialMediaID, userID)
	if err != nil {
		return socialmedia, err
	}
	t := time.Now()

	socialmedia.Name = inputData.Name
	socialmedia.SocialMediaUrl = inputData.SocialMediaUrl
	socialmedia.UpdatedAt = &t

	updatedSocialMedia, err := s.repository.Update(socialmedia)
	if err != nil {
		return updatedSocialMedia, err
	}

	return updatedSocialMedia, nil
}

func (s *service) DeleteSocialMedia(socialMediaID uint, userID uint) error {
	return s.repository.DeleteByID(socialMediaID, userID)
}
