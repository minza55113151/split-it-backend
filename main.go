package main

import (
	"split-it/handlers"
	"split-it/repositories"
	"split-it/services"

	_ "split-it/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {

	app := fiber.New()

	db := InitDB()

	userRepo := repositories.NewUserRepository(db)
	// friendRepo := repositories.NewFriendRepository(db)
	userService := services.NewUserService(userRepo)
	// friendService := services.NewFriendService(friendRepo)
	userHandler := handlers.NewUserHandler(userService)

	app.Get("/users/:uid", userHandler.HandleGetUser)
	app.Post("/users", userHandler.HandleCreateUser)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Listen(":3000")
}
