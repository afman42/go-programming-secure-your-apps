package models

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username"`
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Age      uint   `gorm:"not null" json:"age"`
	// Product  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:set NULL;" json:"products"`
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.Password = helpers.HashPass(u.Password)
// 	return
// }
