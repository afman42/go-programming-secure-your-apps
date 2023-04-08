package photo

import (
	"errors"
	"sesi_4_final_project/models"

	"gorm.io/gorm"
)

type Repository interface {
	Save(photo models.Photo) (models.Photo, error)
	FindByID(ID uint) (models.Photo, error)
	FindByUserID(ID uint, userID uint) (models.Photo, error)
	Update(comment models.Photo) (models.Photo, error)
	FindAll(userID uint) ([]models.Photo, error)
	DeleteByID(ID uint, userID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(photo models.Photo) (models.Photo, error) {
	err := r.db.Create(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) FindByUserID(ID uint, userID uint) (models.Photo, error) {
	var photo models.Photo

	result := r.db.Where("id = ?", ID).Where("user_id = ?", userID).Find(&photo)
	err := result.Error
	if err != nil {
		return photo, err
	}

	count := result.RowsAffected

	if count == 0 {
		return photo, errors.New("cannot find row")
	}

	return photo, nil
}

func (r *repository) FindByID(ID uint) (models.Photo, error) {
	var photo models.Photo

	result := r.db.Where("id = ?", ID).Preload("User").Preload("Comments.User").Find(&photo)
	err := result.Error
	if err != nil {
		return photo, err
	}

	count := result.RowsAffected

	if count == 0 {
		return photo, errors.New("cannot find the row")
	}

	return photo, nil
}

func (r *repository) Update(comment models.Photo) (models.Photo, error) {
	err := r.db.Save(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) FindAll(userID uint) ([]models.Photo, error) {
	var photos []models.Photo

	err := r.db.Where("user_id = ?", userID).Preload("User").Preload("Comments.User").Find(&photos).Error
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (r *repository) DeleteByID(ID uint, userID uint) error {
	var photo models.Photo
	result := r.db.Where("id = ?", ID).Where("user_id = ?", userID).Delete(&photo)
	err := result.Error
	if err != nil {
		return err
	}
	count := result.RowsAffected
	if count == 0 {
		return errors.New("cannot find row")
	}
	return nil
}
