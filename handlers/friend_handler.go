package handlers

import (
	"split-it/models"
	"split-it/services"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gofiber/fiber/v2"
)

type FriendHandler struct {
	friendService *services.FriendService
}

func NewFriendHandler(friendService *services.FriendService) *FriendHandler {
	return &FriendHandler{friendService: friendService}
}

// @Summary Get friends
// @Description Get friends
// @Tags friends
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.FriendResponse
// @Failure 500 {string} string
// @Router /friends [get]
func (h *FriendHandler) HandleGetFriends(c *fiber.Ctx) error {
	usr := c.Locals(models.UserContextKey).(*clerk.User)
	subID := usr.ID

	friends, err := h.friendService.GetFriends(subID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.Status(fiber.StatusOK).JSON(friends)
}

// @Summary Create friend
// @Description Create friend
// @Tags friends
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param subID path string true "Friend subID"
// @Success 201 {string} string
// @Failure 500 {string} string
// @Router /friends/{subID} [post]
func (h *FriendHandler) HandleCreateFriend(c *fiber.Ctx) error {
	usr := c.Locals(models.UserContextKey).(*clerk.User)
	subID := usr.ID
	friendSubID := c.Params("subID")

	if subID == friendSubID {
		return c.Status(fiber.StatusBadRequest).SendString("You cannot be friends with yourself")
	}

	err := h.friendService.CreateFriend(subID, friendSubID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.SendStatus(fiber.StatusCreated)
}

// @Summary Delete friend
// @Description Delete friend
// @Tags friends
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param subID path string true "Friend subID"
// @Success 204 {string} string
// @Failure 500 {string} string
// @Router /friends/{subID} [delete]
func (h *FriendHandler) HandleDeleteFriend(c *fiber.Ctx) error {
	usr := c.Locals(models.UserContextKey).(*clerk.User)
	subID := usr.ID
	friendSubID := c.Params("subID")

	err := h.friendService.DeleteFriend(subID, friendSubID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.SendStatus(fiber.StatusNoContent)
}
