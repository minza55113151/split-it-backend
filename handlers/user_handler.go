package handlers

import (
	"split-it/models"
	"split-it/services"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
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
	usr := c.Locals(models.UserContextKey).(*clerk.User)
	subID := usr.ID

	user, err := h.UserService.GetUserBySubID(subID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// @Summary Get users by name
// @Description Get users by name
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param name path string true "User name"
// @Success 200 {array} models.User
// @Failure 500 {string} string "Internal server error"
// @Router /users/{name} [get]
func (h *UserHandler) GetUserByName(c *fiber.Ctx) error {
	name := c.Params("name")

	users, err := h.UserService.GetUsersByName(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.Status(fiber.StatusOK).JSON(users)
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
	usr := c.Locals(models.UserContextKey).(*clerk.User)
	subID := usr.ID
	name := ""
	if usr.Username != nil {
		name = *usr.Username
	} else if usr.FirstName != nil {
		name = *usr.FirstName
	} else {
		name = "Split-It User " + strings.TrimPrefix(subID, "user_")[0:4]
	}
	email := usr.EmailAddresses[0].EmailAddress
	imageURL := usr.ImageURL

	user, err := h.UserService.CreateUser(subID, name, email, *imageURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.User
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal server error"
// @Router /users [put]
func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	usr := c.Locals(models.UserContextKey).(*clerk.User)
	subID := usr.ID

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.SubID != subID {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}

	res, err := h.UserService.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
