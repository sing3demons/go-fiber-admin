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
	userGroup := v1.Group("users")
	userGroup.Use(middlewares.IsAuthenticated)
	{
		userGroup.Put("/info", controllers.UpdateInfo)
		userGroup.Put("/password", controllers.UpdatePassword)
		userGroup.Get("", controllers.AllUser)
		userGroup.Get("/:id", controllers.GetUser)
		userGroup.Post("", controllers.CreateUser)
		userGroup.Put("/:id", controllers.UpdateUser)
		userGroup.Delete("/:id", controllers.DeleteUser)

	}

	roleGroup := v1.Group("roles/")
	roleGroup.Use(middlewares.IsAuthenticated)
	{

		roleGroup.Get("", controllers.AllRoles)
		roleGroup.Get("/:id", controllers.GetRole)
		roleGroup.Post("", controllers.CreateRole)
		roleGroup.Put("/:id", controllers.UpdateRole)
		roleGroup.Delete("/:id", controllers.DeleteRole)

	}

	permissionGroup := v1.Group("permission/")
	permissionGroup.Use(middlewares.IsAuthenticated)
	permissionGroup.Get("", controllers.AllPermissions)

	productGroup := v1.Group("products/")
	productGroup.Use(middlewares.IsAuthenticated)
	{
		productGroup.Get("", controllers.AllProducts)
		productGroup.Get("/:id", controllers.GetProduct)
		productGroup.Post("", controllers.CreateProduct)
		productGroup.Post("/uploads", controllers.UploadImage)
		productGroup.Put("/:id", controllers.UpdateProduct)
		productGroup.Delete("/:id", controllers.DeleteProduct)

	}

	orderGroup := v1.Group("orders/")
	orderGroup.Use(middlewares.IsAuthenticated)
	{
		orderGroup.Get("", controllers.AllOrders)
		orderGroup.Post("/export", controllers.Export)
		orderGroup.Get("/chart", controllers.Chart)
	}

}
