package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Email     string `gorm:"unique" json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"`
	RoleID    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleID"`
}

func (user *User) EncryptedPassword(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(hash)
}

func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}

func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User
	db.Offset(offset).Limit(limit).Find(&users)
	return users
}
