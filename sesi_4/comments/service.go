package comments

import (
	"sesi_4_final_project/models"
	"time"
)

type Service interface {
	CreateComment(input CreateCommentInput, userID uint, photoID uint) (models.Comment, error)
	GetAll(photoID uint) ([]models.Comment, error)
	GetOne(commentID uint, photoID uint) (models.Comment, error)
	UpdateComment(commentID uint, input CreateCommentInput, userID uint, photoID uint) (models.Comment, error)
	DeleteComment(commentID uint, userID uint, photoID uint) error
	GetOneByUserID(commentID uint, userID uint, photoID uint) (models.Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateComment(input CreateCommentInput, userID uint, photoID uint) (models.Comment, error) {
	comment := models.Comment{}
	comment.Message = input.Message
	comment.UserID = userID
	comment.PhotoID = photoID

	newComment, err := s.repository.Save(comment)
	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

func (s *service) GetAll(photoID uint) ([]models.Comment, error) {
	comments, err := s.repository.FindAll(photoID)
	if err != nil {
		return []models.Comment{}, err
	}

	return comments, nil
}

func (s *service) GetOneByUserID(commentID uint, userID uint, photoID uint) (models.Comment, error) {
	photo, err := s.repository.FindByUserID(commentID, userID, photoID)
	if err != nil {
		return models.Comment{}, err
	}

	return photo, nil
}

func (s *service) GetOne(commentID uint, photoID uint) (models.Comment, error) {
	photo, err := s.repository.FindByID(commentID, photoID)
	if err != nil {
		return models.Comment{}, err
	}

	return photo, nil
}

func (s *service) UpdateComment(commentID uint, inputData CreateCommentInput, userID uint, photoID uint) (models.Comment, error) {
	comment, err := s.repository.FindByUserID(commentID, userID, photoID)
	if err != nil {
		return comment, err
	}
	t := time.Now()

	comment.Message = inputData.Message
	comment.UpdatedAt = &t

	updatedComment, err := s.repository.Update(comment)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}

func (s *service) DeleteComment(commentID uint, userID uint, photoID uint) error {
	return s.repository.DeleteByID(commentID, userID, photoID)
}
