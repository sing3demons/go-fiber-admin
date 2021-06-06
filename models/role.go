package models

type Role struct {
	ID         uint         `gorm:"primarykey" json:"id"`
	Name       string       `json:"name"`
	Permission []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}
