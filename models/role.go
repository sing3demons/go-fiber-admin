package models

type Role struct {
	ID         uint         `gorm:"primarykey" json:"id"`
	Name       string       `json:"name"`
	Permission []Permission `json:"permission" gorm:"many2many:role_permissions"`
}
