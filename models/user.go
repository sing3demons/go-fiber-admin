package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Email     string `gorm:"unique" json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"`

}

func (user *User) EncryptedPassword(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(hash)
}
