package models

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex"`
	Email    string `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`
	Age      uint   `gorm:"not null"`
	// Product  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:set NULL;" json:"products"`
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.Password = helpers.HashPass(u.Password)
// 	return
// }
