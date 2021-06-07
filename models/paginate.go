package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, page int, limit int, entity Entity) fiber.Map {
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)

	total := entity.Count(db)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	}

}
