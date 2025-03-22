package main

import (
	"split-it/handlers"
	"split-it/repositories"
	"split-it/services"

	_ "split-it/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Split-It API
// @description This is a sample server for the Split-It application.
// @version 1.0
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	app := fiber.New()

	db := initDB()
	initAuth()

	userRepo := repositories.NewUserRepository(db)
	// friendRepo := repositories.NewFriendRepository(db)
	userService := services.NewUserService(userRepo)
	// friendService := services.NewFriendService(friendRepo)
	userHandler := handlers.NewUserHandler(userService)

	app.Get("/swagger/*", swagger.New(swagger.Config{
		PersistAuthorization: true,
	}))

	authGroup := app.Group("")
	authGroup.Use(authMiddleware)
	authGroup.Get("/users", userHandler.HandleGetUser)
	authGroup.Post("/users", userHandler.HandleCreateUser)

	app.Listen(":8000")
}
