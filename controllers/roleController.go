package controllers

import (
	"github/sing3demons/go-fiber-admin/database"
	"github/sing3demons/go-fiber-admin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RoleCreateDTO struct {
	name        string
	permissions []string
}

func AllRoles(c *fiber.Ctx) error {
	var role []models.Role
	database.DB.Preload("Permission").Find(&role)
	return c.JSON(role)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}
	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			ID: uint(id),
		}
	}

	role := models.Role{
		Name:       roleDto["name"].(string),
		Permission: permissions,
	}

	if err := database.DB.Create(&role).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(role)

}

func GetRole(c *fiber.Ctx) error {
	user, err := findRoleByID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {

		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			ID: uint(id),
		}

	}

	var result interface{}

	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		ID:         uint(id),
		Name:       roleDto["name"].(string),
		Permission: permissions,
	}

	database.DB.Model(&role).Updates(role)

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

func findRoleByID(c *fiber.Ctx) (*models.Role, error) {
	id, _ := strconv.Atoi(c.Params("id"))
	var role models.Role
	role.ID = uint(id)

	if err := database.DB.First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}
