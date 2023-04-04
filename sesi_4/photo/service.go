package photo

import (
	"sesi_4_final_project/models"
	"time"
)

type Service interface {
	CreatePhoto(input CreatePhotoInput, userID uint) (models.Photo, error)
	GetAll(userID uint) ([]models.Photo, error)
	GetOne(photoID uint) (models.Photo, error)
	UpdatePhoto(photoID uint, input CreatePhotoInput, userID uint) (models.Photo, error)
	DeletePhoto(photoID uint, userID uint) error
	GetOneByUserID(photoID uint, userID uint) (models.Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePhoto(input CreatePhotoInput, userID uint) (models.Photo, error) {
	photo := models.Photo{}
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl
	photo.UserID = userID

	newPhoto, err := s.repository.Save(photo)
	if err != nil {
		return newPhoto, err
	}

	return newPhoto, nil
}

func (s *service) GetAll(userID uint) ([]models.Photo, error) {
	photos, err := s.repository.FindAll(userID)
	if err != nil {
		return []models.Photo{}, err
	}

	return photos, nil
}

func (s *service) GetOneByUserID(photoID uint, userID uint) (models.Photo, error) {
	photo, err := s.repository.FindByUserID(photoID, userID)
	if err != nil {
		return models.Photo{}, err
	}

	return photo, nil
}

func (s *service) GetOne(photoID uint) (models.Photo, error) {
	photo, err := s.repository.FindByID(photoID)
	if err != nil {
		return models.Photo{}, err
	}

	return photo, nil
}

func (s *service) UpdatePhoto(photoID uint, inputData CreatePhotoInput, userID uint) (models.Photo, error) {
	photo, err := s.repository.FindByUserID(photoID, userID)
	if err != nil {
		return photo, err
	}
	t := time.Now()

	photo.Title = inputData.Title
	photo.Caption = inputData.Caption
	photo.PhotoUrl = inputData.PhotoUrl
	photo.UpdatedAt = &t

	updatedphoto, err := s.repository.Update(photo)
	if err != nil {
		return updatedphoto, err
	}

	return updatedphoto, nil
}

func (s *service) DeletePhoto(photoID uint, userID uint) error {
	return s.repository.DeleteByID(photoID, userID)
}
