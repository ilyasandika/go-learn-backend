package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"time"
	"uaspw2/controller"
	"uaspw2/repository"
	"uaspw2/routes"
	"uaspw2/services"
)

func main() {
	app := fiber.New()
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/uaspw2")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	userRepository := repository.NewUserRepository()
	userServices := services.NewUserService(userRepository, db)
	userController := controller.NewUserController(userServices)
	routes.SetupUserRoutes(app, userController)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Log message when the server starts successfully
	log.Info("Server is running on port 3000")
	select {}
}
