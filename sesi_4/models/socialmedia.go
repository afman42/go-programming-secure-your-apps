package models

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null;varchar(191)" json:"name"`
	SocialMediaUrl string `gorm:"not null;varchar(191)" json:"photo_id"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	User           *User  `json:"user"`
}
