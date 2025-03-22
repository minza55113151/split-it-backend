package handlers

import (
	"split-it/models"
	"split-it/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// @Summary Get a user
// @Description Get user details
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.User
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users [get]
func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	uid := c.Locals(models.UserIdContextKey).(string)

	user, err := h.UserService.GetUserByUID(uid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 201 {object} models.User
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	uid := c.Locals(models.UserIdContextKey).(string)

	user, err := h.UserService.CreateUser(uid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
