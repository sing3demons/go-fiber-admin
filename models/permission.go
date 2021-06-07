package models

type Permission struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}


