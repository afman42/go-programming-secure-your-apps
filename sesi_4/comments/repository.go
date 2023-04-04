package comments

import (
	"errors"
	"sesi_4_final_project/models"

	"gorm.io/gorm"
)

type Repository interface {
	Save(comment models.Comment) (models.Comment, error)
	FindByID(ID uint, photoID uint) (models.Comment, error)
	FindByUserID(ID uint, userID uint, photoID uint) (models.Comment, error)
	Update(comment models.Comment) (models.Comment, error)
	FindAll(photoID uint) ([]models.Comment, error)
	DeleteByID(ID uint, userID uint, photoID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(comment models.Comment) (models.Comment, error) {
	err := r.db.Create(&comment).Error
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) FindByUserID(ID uint, userID uint, photoID uint) (models.Comment, error) {
	var comment models.Comment

	result := r.db.Where("id = ?", ID).Where("photo_id = ?", userID).Where("user_id = ?", userID).Find(&comment)
	err := result.Error
	if err != nil {
		return comment, err
	}

	count := result.RowsAffected

	if count == 0 {
		return comment, errors.New("cannot find row comment")
	}

	return comment, nil
}

func (r *repository) FindByID(ID uint, photoID uint) (models.Comment, error) {
	var comment models.Comment

	result := r.db.Where("id = ?", ID).Where("photo_id = ?", photoID).Find(&comment)
	err := result.Error
	if err != nil {
		return comment, err
	}

	count := result.RowsAffected

	if count == 0 {
		return comment, errors.New("cannot find row comment")
	}

	return comment, nil
}

func (r *repository) Update(comment models.Comment) (models.Comment, error) {
	err := r.db.Save(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) FindAll(photoID uint) ([]models.Comment, error) {
	var comments []models.Comment

	err := r.db.Where("photo_id = ?", photoID).Find(&comments).Error
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (r *repository) DeleteByID(ID uint, userID uint, photoID uint) error {
	var comment models.Comment
	result := r.db.Where("photo_id = ?", photoID).Where("id = ?", ID).Where("user_id = ?", userID).Delete(&comment)
	err := result.Error
	if err != nil {
		return err
	}
	count := result.RowsAffected
	if count == 0 {
		return errors.New("cannot find row comment")
	}
	return nil
}
