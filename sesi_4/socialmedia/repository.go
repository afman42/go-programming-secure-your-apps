package socialmedia

import (
	"fmt"
	"sesi_4_final_project/models"

	"gorm.io/gorm"
)

type Repository interface {
	Save(socialmedia models.SocialMedia) (models.SocialMedia, error)
	FindByID(ID uint) (models.SocialMedia, error)
	FindByUserID(ID uint, userID uint) (models.SocialMedia, error)
	Update(socialmedia models.SocialMedia) (models.SocialMedia, error)
	FindAll(userID uint) ([]models.SocialMedia, error)
	DeleteByID(ID uint, userID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(socialmedia models.SocialMedia) (models.SocialMedia, error) {
	err := r.db.Create(&socialmedia).Error
	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r *repository) FindByUserID(ID uint, userID uint) (models.SocialMedia, error) {
	var socialmedia models.SocialMedia

	result := r.db.Where("id = ?", ID).Where("user_id = ?", userID).Find(&socialmedia)
	err := result.Error
	if err != nil {
		return socialmedia, err
	}

	count := result.RowsAffected

	if count == 0 {
		return socialmedia, fmt.Errorf("cannot find id %d social media", ID)
	}

	return socialmedia, nil
}

func (r *repository) FindByID(ID uint) (models.SocialMedia, error) {
	var socialmedia models.SocialMedia

	result := r.db.Where("id = ?", ID).Find(&socialmedia)
	err := result.Error
	if err != nil {
		return socialmedia, err
	}

	count := result.RowsAffected

	if count == 0 {
		return socialmedia, fmt.Errorf("cannot find id %d social media", ID)
	}

	return socialmedia, nil
}

func (r *repository) Update(socialmedia models.SocialMedia) (models.SocialMedia, error) {
	err := r.db.Save(&socialmedia).Error

	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r *repository) FindAll(userID uint) ([]models.SocialMedia, error) {
	var socialmedias []models.SocialMedia

	err := r.db.Where("user_id = ?", userID).Find(&socialmedias).Error
	if err != nil {
		return socialmedias, err
	}

	return socialmedias, nil
}

func (r *repository) DeleteByID(ID uint, userID uint) error {
	var socialmedia models.SocialMedia
	result := r.db.Where("id = ?", ID).Where("user_id = ?", userID).Delete(&socialmedia)
	err := result.Error
	if err != nil {
		return err
	}
	count := result.RowsAffected
	if count == 0 {
		return fmt.Errorf("cannot find id %d social media", ID)
	}
	return nil
}
