package routes

import (
	"github/sing3demons/go-fiber-admin/controllers"

	"github.com/gofiber/fiber/v2"
)

func Serve(app *fiber.App) {
	v1 := app.Group("api/v1/")

	authGroup := v1.Group("auth/")
	{
		authGroup.Post("register", controllers.Register)
		authGroup.Post("login", controllers.Login)
	}
}
