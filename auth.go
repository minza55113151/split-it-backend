package main

import (
	"fmt"
	"log"
	"os"
	"split-it/models"
	"strings"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func initAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))
}

func authMiddleware(c *fiber.Ctx) error {
	sessionToken := c.Get("Authorization")
	sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")

	claims, err := jwt.Verify(c.Context(), &jwt.VerifyParams{
		Token:  sessionToken,
		Leeway: time.Hour * 24 * 100000,
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	usr, err := user.Get(c.Context(), claims.Subject)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if usr.ID == "" {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error - User ID not found")
	}

	fmt.Println(usr.ID) // TODO: remove this
	c.Locals(models.SubIDContextKey, usr.ID)

	return c.Next()
}
