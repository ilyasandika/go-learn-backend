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
		userGroup.Get("/details", middlewares.AuthRequired, controller.FindByToken)
		userGroup.Get("/:userId", middlewares.AdminOnly, controller.FindByPath)
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
		authGroup.Post("/verify-auth", controller.VerifyAuth)
	}
}

func SetupUserProfileRoutes(app *fiber.App, controller controllers.UserProfileController) {
	apiGroup := app.Group("/api")
	userProfileGroup := apiGroup.Group("/user_profiles")
	{
		userProfileGroup.Get("/", middlewares.AdminOnly, controller.FindAll)
		userProfileGroup.Get("/details", middlewares.AuthRequired, controller.FindByToken)
		userProfileGroup.Put("/details", middlewares.AuthRequired, controller.UpdateByToken)
		userProfileGroup.Get("/details/:userId", middlewares.AuthRequired, controller.FindByPath)

	}
}

func SetupUserProfilePhotoRoutes(app *fiber.App, controller controllers.UserProfilePhotoController) {
	apiGroup := app.Group("/api")
	userProfilePhoto := apiGroup.Group("/user_profile_photos")
	{
		userProfilePhoto.Get("/", middlewares.AuthRequired, controller.FindByToken)
		userProfilePhoto.Put("/", middlewares.AuthRequired, controller.UpdateByToken)
	}
}

func SetupArticlePhotoRoutes(app *fiber.App, controller controllers.ArticleController) {
	apiGroup := app.Group("/api")
	articleGroup := apiGroup.Group("/articles")
	{
		articleGroup.Post("/", middlewares.AuthRequired, controller.CreateByToken)
		articleGroup.Get("/published", middlewares.AuthRequired, controller.FindAllPublished)
		articleGroup.Put("/published/:articleId", middlewares.AdminOnly, controller.PublishArticle)
		articleGroup.Get("/published/user", middlewares.UserOnly, controller.FindAllPublishedByUserID)
		articleGroup.Get("/unpublished", middlewares.AdminOnly, controller.FindAllUnpublished)
		articleGroup.Get("/unpublished/user", middlewares.UserOnly, controller.FindAllUnpublishedByUserID)
		articleGroup.Put("/unpublished/:articleId", middlewares.AdminOnly, controller.UnpublishArticle)
		articleGroup.Get("/:articleId", middlewares.AuthRequired, controller.FindByID)
		articleGroup.Put("/:articleId", middlewares.UserOnly, controller.UpdateByID)
		articleGroup.Delete("/:articleId", middlewares.AuthRequired, controller.DeleteByID)
	}
}

func SetupLikeRoutes(app *fiber.App, controller controllers.LikeController) {
	apiGroup := app.Group("/api")
	likeGroup := apiGroup.Group("/likes")
	{
		likeGroup.Get("/articles/:articleId", middlewares.AuthRequired, controller.FindByArticleId)
		likeGroup.Post("/articles/:articleId", middlewares.UserOnly, controller.Create)
		likeGroup.Delete("/articles/:articleId", middlewares.UserOnly, controller.Delete)
		likeGroup.Get("/users/:userId", middlewares.AuthRequired, controller.FindByUserId)
	}
}

func SetupCommentRoutes(app *fiber.App, controller controllers.CommentController) {
	apiGroup := app.Group("/api")
	likeGroup := apiGroup.Group("/comments")
	{
		likeGroup.Get("/articles/:articleId", middlewares.AuthRequired, controller.FindByArticleId)
		likeGroup.Post("/articles/:articleId", middlewares.UserOnly, controller.Create)
		likeGroup.Delete("/:commentId", middlewares.UserOnly, controller.Delete)
	}
}
