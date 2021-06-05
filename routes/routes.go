package routes

import (
	"github/sing3demons/go-fiber-admin/controllers"
	"github/sing3demons/go-fiber-admin/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Serve(app *fiber.App) {
	v1 := app.Group("api/v1/")

	authGroup := v1.Group("auth/")
	{
		authGroup.Post("register", controllers.Register)
		authGroup.Post("login", controllers.Login)
		authGroup.Use(middlewares.IsAuthenticated)
		authGroup.Get("user", controllers.User)
		authGroup.Get("logout", controllers.Logout)
	}

}
