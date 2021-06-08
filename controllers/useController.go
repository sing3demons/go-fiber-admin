package controllers

import (
	"github/sing3demons/go-fiber-admin/database"
	"github/sing3demons/go-fiber-admin/middlewares"
	"github/sing3demons/go-fiber-admin/models"
	"github/sing3demons/go-fiber-admin/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type formCreate struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RoleID    string `json:"role_id"`
}

func AllUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "5"))

	return c.JSON(models.Paginate(database.DB.Preload("Role"), page, limit, &models.User{}))
}

func CreateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var form formCreate
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	rId, _ := strconv.Atoi(form.RoleID)
	var user models.User

	copier.Copy(&user, &form)
	user.RoleID = uint(rId)
	user.EncryptedPassword("12345678")
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
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	user, err := findUserByID(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	// user := models.User{
	// 	ID: uint(id),
	// }
	var form formCreate
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	rId, _ := strconv.Atoi(form.RoleID)
	var user models.User

	copier.Copy(&user, &form)
	user.ID = uint(id)
	user.RoleID = uint(rId)

	database.DB.Model(&user).Updates(user)

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
