package models

import (
	"sesi_2_authentication_middleware/enums"
	"sesi_2_authentication_middleware/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Email    string         `gorm:"not null;uniqueIndex" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string         `gorm:"not null" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role     enums.RoleUser `gorm:"type:enum('admin','user')" form:"role" valid:"required~Your role is required"`
}

type LoginUser struct {
	Email    string `form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
