package models

type User struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Email     string `gorm:"unique" json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"`
}
