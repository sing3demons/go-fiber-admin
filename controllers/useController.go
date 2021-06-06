package controllers

import (
	"github/sing3demons/go-fiber-admin/database"
	"github/sing3demons/go-fiber-admin/models"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllUser(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "5"))
	offset := (page - 1) * limit
	var total int64

	users := []models.User{}
	database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	database.DB.Model(&models.User{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.EncryptedPassword(user.Password)
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)

}

func findUserByID(c *fiber.Ctx) (*models.User, error) {
	id, _ := strconv.Atoi(c.Params("id"))
	var user models.User
	user.ID = uint(id)

	if err := database.DB.Preload("Role").First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(c *fiber.Ctx) error {
	user, err := findUserByID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	user, err := findUserByID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	var form models.User
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	database.DB.Model(&user).Updates(form)

	return c.SendStatus(fiber.StatusNoContent)

}

func DeleteUser(c *fiber.Ctx) error {
	user, err := findUserByID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Delete(&user)

	return c.SendStatus(fiber.StatusNoContent)

}
