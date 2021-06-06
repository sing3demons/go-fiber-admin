package controllers

import (
	"github/sing3demons/go-fiber-admin/database"
	"github/sing3demons/go-fiber-admin/models"
	"github/sing3demons/go-fiber-admin/util"
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

	lastPage := math.Ceil(float64(int(total) / limit))

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
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

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)

	user := models.User{
		ID:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)
	user := models.User{ID: uint(userId)}
	user.EncryptedPassword(data["password"])

	database.DB.Model(&user).Updates(user)

	return c.SendStatus(fiber.StatusOK)
}
