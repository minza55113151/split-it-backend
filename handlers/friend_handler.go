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

func (h *FriendHandler) HandleGetFriends(c *fiber.Ctx) error {
	subID := c.Locals(models.SubIDContextKey).(string)

	friends, err := h.FriendService.GetFriends(subID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.Status(fiber.StatusOK).JSON(friends)
}

func (h *FriendHandler) HandleCreateFriend(c *fiber.Ctx) error {
	subID := c.Locals(models.SubIDContextKey).(string)
	friendSubID := c.Params("subID")

	err := h.FriendService.CreateFriend(subID, friendSubID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *FriendHandler) HandleDeleteFriend(c *fiber.Ctx) error {
	subID := c.Locals(models.SubIDContextKey).(string)
	friendSubID := c.Params("subID")

	err := h.FriendService.DeleteFriend(subID, friendSubID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.SendStatus(fiber.StatusNoContent)
}
