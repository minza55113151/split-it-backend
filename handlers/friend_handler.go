package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type FriendHandler struct{}

func NewFriendHandler() *FriendHandler {
	return &FriendHandler{}
}

func (h *FriendHandler) AddFriend(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
