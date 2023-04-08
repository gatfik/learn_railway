package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required~title is required"`
	Description string `json:"description" valid:"required~description is required"`
	UserId      uint   `json:"user_id"`
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return
	}

	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return
	}

	return
}
