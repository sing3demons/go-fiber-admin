package controllers

import (
	"github/sing3demons/go-fiber-admin/database"
	"github/sing3demons/go-fiber-admin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var role []models.Role
	database.DB.Find(&role)
	return c.JSON(role)
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return err
	}

	if err := database.DB.Create(&role).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(role)

}

func findRoleByID(c *fiber.Ctx) (*models.Role, error) {
	id, _ := strconv.Atoi(c.Params("id"))
	var role models.Role
	role.ID = uint(id)

	if err := database.DB.First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func GetRole(c *fiber.Ctx) error {
	user, err := findRoleByID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateRole(c *fiber.Ctx) error {
	role, err := findRoleByID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	var form models.Role
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	database.DB.Model(&role).Updates(form)

	return c.JSON(role)

}

func DeleteRole(c *fiber.Ctx) error {
	role, err := findRoleByID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Delete(&role)

	return c.SendStatus(fiber.StatusNoContent)

}
