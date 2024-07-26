package routes

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/controllers"
	"uaspw2/middlewares"
)

func SetupUserRoutes(app *fiber.App, controller controllers.UserController) {
	apiGroup := app.Group("/api")
	userGroup := apiGroup.Group("/users")
	{
		userGroup.Get("/", middlewares.AdminOnly, controller.FindAll)
		userGroup.Get("/:userId", controller.FindById)
		userGroup.Put("/:userId", middlewares.AdminOnly, controller.UpdateByPath)
		userGroup.Put("/", middlewares.AuthRequired, controller.UpdateByToken)
		userGroup.Delete("/:userId", middlewares.AdminOnly, controller.Delete)
	}
}

func SetupAuthRoutes(app *fiber.App, controller controllers.AuthController) {
	apiGroup := app.Group("/api")
	authGroup := apiGroup.Group("/auth")
	{
		authGroup.Post("/login", middlewares.GuestOnly, controller.Login)
		authGroup.Post("/logout", middlewares.AuthRequired, controller.Logout)
		authGroup.Post("/register", middlewares.GuestOnly, controller.Register)
	}
}
