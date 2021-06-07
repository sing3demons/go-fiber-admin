package controllers

import (
	"github/sing3demons/go-fiber-admin/database"
	"github/sing3demons/go-fiber-admin/models"
	"github/sing3demons/go-fiber-admin/util"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleID:    2,
	}
	user.EncryptedPassword(data["password"])

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	if err := database.DB.Where("email = ?", data["email"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email or password incorect"})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.SendStatus(fiber.StatusNoContent)
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	var user models.User
	database.DB.Where("id = ?", id).Preload("Role").First(&user)
	return c.Status(fiber.StatusOK).JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.SendStatus(fiber.StatusNoContent)
}
