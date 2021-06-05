package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	FirstName string
	LastName  string
	Password  string
}
