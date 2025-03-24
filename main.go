package main

import (
	"split-it/handlers"
	"split-it/repositories"
	"split-it/services"

	_ "split-it/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title Split-It API
// @description This is a sample server for the Split-It application.
// @version 1.0
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	app := fiber.New()

	app.Use(cors.New())

	db := initDB()
	initAuth()

	userRepo := repositories.NewUserRepository(db)
	friendRepo := repositories.NewFriendRepository(db)
	expenseRepo := repositories.NewExpenseRepository(db)

	userService := services.NewUserService(userRepo)
	friendService := services.NewFriendService(friendRepo)
	expenseService := services.NewExpenseService(expenseRepo)

	userHandler := handlers.NewUserHandler(userService)
	friendHandler := handlers.NewFriendHandler(friendService)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

	app.Get("/swagger/*", swagger.New(swagger.Config{
		PersistAuthorization: true,
	}))

	authGroup := app.Group("")
	authGroup.Use(authMiddleware)

	authGroup.Get("/users", userHandler.HandleGetUser)
	authGroup.Post("/users", userHandler.HandleCreateUser)

	authGroup.Get("/friends", friendHandler.HandleGetFriends)
	authGroup.Post("/friends/:subID", friendHandler.HandleCreateFriend)
	authGroup.Delete("/friends/:subID", friendHandler.HandleDeleteFriend)

	authGroup.Get("/expenses/:status", expenseHandler.HandleGetUserExpensesWithStatus)
	authGroup.Post("/expenses", expenseHandler.HandleCreateExpense)
	authGroup.Put("/expenses/:id", expenseHandler.HandleUpdateExpense)
	authGroup.Delete("/expenses/:id", expenseHandler.HandleDeleteExpense)

	app.Listen(":8000")
}
