package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint        `json:"id"`
	Email      string      `json:"email"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Name       string      `json:"name" gorm:"-"`
	Total      float32     `json:"total" gorm:"-"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID           uint    `json:"id"`
	OrderID      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(order).Count(&total)

	return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order

	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	for i, _ := range orders {
		var total float32 = 0

		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float32(orderItem.Quantity)
		}
		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
		orders[i].Total = total
	}

	return orders
}
