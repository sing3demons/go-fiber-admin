package models

import "gorm.io/gorm"

type Product struct {
	ID          uint    `gorm:"primarykey" json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

func (product *Product) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Product{}).Count(&total)
	return total
}

func (product *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []Product
	db.Offset(offset).Limit(limit).Order("id desc").Find(&products)
	return products
}
