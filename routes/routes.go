package routes

import (
	"github/sing3demons/go-fiber-admin/controllers"

	"github.com/gofiber/fiber/v2"
)

func Serve(app *fiber.App) {
	app.Get("", controllers.Register)
}
