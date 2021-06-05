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
		authGroup.Get("/user", controllers.User)
		authGroup.Get("logout", controllers.Logout)
	}
	userGroup := v1.Group("users/")
	userGroup.Use(middlewares.IsAuthenticated)
	{

		userGroup.Get("", controllers.AllUser)
		userGroup.Get("/:id", controllers.GetUser)
		userGroup.Post("", controllers.CreateUser)
		userGroup.Post("/:id", controllers.UpdateUser)
		userGroup.Delete("/:id", controllers.DeleteUser)

	}

}
