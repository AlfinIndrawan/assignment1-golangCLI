package models

import (
	"jwt-assignment/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" validation:"required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" validation:"required"`
	Password string    `gorm:"not null" json:"password" validation:"required,min=6"`
	Products []Product `gorm:"constraint:onUpdate:CASCADE,onDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		return errCreate
	}
	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}
