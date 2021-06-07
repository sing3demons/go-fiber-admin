package middlewares

import (
	"errors"
	"fmt"
	"github/sing3demons/go-fiber-admin/database"
	"github/sing3demons/go-fiber-admin/models"
	"github/sing3demons/go-fiber-admin/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")

	Id, err := util.ParseJwt(cookie)

	fmt.Println(Id)

	if err != nil {
		return err
	}

	userId, _ := strconv.Atoi(Id)

	user := models.User{
		ID: uint(userId),
	}

	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		ID: user.RoleID,
	}

	database.DB.Preload("Permission").Find(&role)

	fmt.Println(role.Permission)
	if c.Method() == fiber.MethodGet {
		for _, permission := range role.Permission {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permission {
			fmt.Println(permission.Name)
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}
