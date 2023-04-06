package models

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex;unique" json:"username"`
	Email    string `gorm:"not null;uniqueIndex;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Age      uint   `gorm:"not null" json:"age"`
}
