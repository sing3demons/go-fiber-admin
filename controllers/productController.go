package controllers

import (
	"github/sing3demons/go-fiber-admin/database"
	"github/sing3demons/go-fiber-admin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "5"))

	return c.JSON(models.Paginate(database.DB, page, limit, &models.Product{}))
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)

}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{
		ID: uint(id),
	}
	database.DB.Find(&product)
	return c.Status(fiber.StatusOK).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	product, _ := findProductByID(c)

	var form models.Product

	if err := c.BodyParser(&form); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(form)
	return c.Status(fiber.StatusOK).JSON(product)
}

func findProductByID(c *fiber.Ctx) (*models.Product, error) {
	id, _ := strconv.Atoi(c.Params("id"))
	var product models.Product
	product.ID = uint(id)

	if err := database.DB.First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{
		ID: uint(id),
	}

	database.DB.Delete(&product)
	return c.SendStatus(fiber.StatusOK)
}
