package handlers

import (
	"split-it/models"
	"split-it/services"

	"github.com/gofiber/fiber/v2"
)

type FriendHandler struct {
	FriendService *services.FriendService
}

func NewFriendHandler(friendService *services.FriendService) *FriendHandler {
	return &FriendHandler{FriendService: friendService}
}

// @Summary Get friends
// @Description Get friends
// @Tags friends
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Friend
// @Failure 500 {string} string
// @Router /friends [get]
func (h *FriendHandler) HandleGetFriends(c *fiber.Ctx) error {
	subID := c.Locals(models.SubIDContextKey).(string)

	friends, err := h.FriendService.GetFriends(subID)
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
	subID := c.Locals(models.SubIDContextKey).(string)
	friendSubID := c.Params("subID")

	err := h.FriendService.CreateFriend(subID, friendSubID)
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
	subID := c.Locals(models.SubIDContextKey).(string)
	friendSubID := c.Params("subID")

	err := h.FriendService.DeleteFriend(subID, friendSubID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.SendStatus(fiber.StatusNoContent)
}
