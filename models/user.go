package models

import (
	"dts/learn_middleware/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `json:"full_name" gorm:"not null" valid:"required~Your full name is required"`
	Email    string    `json:"email" gorm:"not null,uniqueIndex" form:"email" valid:"required~Your email is required,email~invalid email format"`
	Password string    `json:"password" gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password minimal 6 character"`
	Products []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	u.Password, _ = helpers.HashPassword(u.Password)
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	return
}
