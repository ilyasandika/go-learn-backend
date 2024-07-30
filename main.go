package main

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"time"
	"uaspw2/controllers"
	"uaspw2/exception"
	"uaspw2/repositories"
	"uaspw2/routes"
	"uaspw2/services"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/uaspw2")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	validate := validator.New()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db, validate)
	userController := controllers.NewUserController(userService)

	userProfileRepository := repositories.NewUserProfileRepository()
	userProfileService := services.NewUserProfileService(userProfileRepository, db, validate)
	userProfileController := controllers.NewUserProfileController(userProfileService)

	authRepository := repositories.NewAuthenticationRepository()
	authService := services.NewAuthenticationServices(authRepository, db, validate)
	authController := controllers.NewAuthenticationController(authService)

	userProfilePhotoRepository := repositories.NewUserProfilePhotoRepository()
	userProfilePhotoService := services.NewUserProfilePhotoService(userProfilePhotoRepository, db, validate)
	userProfilePhotoController := controllers.NewUserProfilePhotoController(userProfilePhotoService)

	articleRepository := repositories.NewArticleRepository()
	articleService := services.NewArticleService(articleRepository, db, validate)
	articleController := controllers.NewArticleController(articleService)

	likeRepository := repositories.NewLikeRepository()
	likeService := services.NewLikeService(likeRepository, db, validate)
	likeController := controllers.NewLikeController(likeService)

	commentRepository := repositories.NewCommentRepository()
	commentService := services.NewCommentService(commentRepository, db, validate)
	commentController := controllers.NewCommentController(commentService)

	app.Use(recover.New())
	app.Static("/", "./public")

	routes.SetupUserRoutes(app, userController)
	routes.SetupUserProfileRoutes(app, userProfileController)
	routes.SetupUserProfilePhotoRoutes(app, userProfilePhotoController)
	routes.SetupAuthRoutes(app, authController)
	routes.SetupArticlePhotoRoutes(app, articleController)
	routes.SetupLikeRoutes(app, likeController)
	routes.SetupCommentRoutes(app, commentController)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Info("Server is running on port 3000")
	select {}
}
