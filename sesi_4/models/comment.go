package models

type Comment struct {
	GormModel
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Message string `gorm:"not null" json:"message"`
	// Product  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:set NULL;" json:"products"`
}
