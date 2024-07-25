package routes

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/controllers"
)

func SetupUserRoutes(app *fiber.App, controller controllers.UserController) {
	apiGroup := app.Group("/api")
	userGroup := apiGroup.Group("/users")
	{
		userGroup.Get("/", controller.FindAll)
		userGroup.Get("/:userId", controller.FindById)
		userGroup.Post("/", controller.Create)
		userGroup.Put("/:userId", controller.Update)
		userGroup.Delete("/:userId", controller.Delete)
	}
}

func SetupAuthRoutes(app *fiber.App, controller controllers.AuthController) {
	apiGroup := app.Group("/api")
	authGroup := apiGroup.Group("/auth")
	{
		authGroup.Post("/login", controller.Login)
		authGroup.Post("/logout", controller.Logout)
	}
}
