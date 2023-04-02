package socialmedia

import (
	"sesi_4_final_project/models"

	"gorm.io/gorm"
)

type Repository interface {
	Save(socialmedia models.SocialMedia) (models.SocialMedia, error)
	FindByID(ID uint) (models.SocialMedia, error)
	Update(socialmedia models.SocialMedia) (models.SocialMedia, error)
	FindAll(userID uint) ([]models.SocialMedia, error)
	DeleteByID(ID uint) error
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

func (r *repository) FindByID(ID uint) (models.SocialMedia, error) {
	var socialmedia models.SocialMedia

	err := r.db.Where("id = ?", ID).Find(&socialmedia).Error
	if err != nil {
		return socialmedia, err
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

func (r *repository) DeleteByID(ID uint) error {
	var socialmedia models.SocialMedia
	r.db.Where("id = ?", ID).Delete(&socialmedia)
	return nil
}
