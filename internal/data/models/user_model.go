package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name             string
	Email            string `gorm:"unique"`
	Age              int32
	CurrentAddress   string `gorm:"unique"`
	PermanentAddress string `gorm:"unique"`
}
