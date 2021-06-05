package controllers

import (
	"github/sing3demons/go-fiber-admin/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "k",
	}

	return c.JSON(user)
}
