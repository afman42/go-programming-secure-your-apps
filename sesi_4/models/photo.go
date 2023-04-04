package models

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	Caption  string `gorm:"null" json:"caption"`
	PhotoUrl string `gorm:"not null;varchar(191)" json:"photo_url"`
}
