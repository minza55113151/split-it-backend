package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type FriendHandler struct{}

func NewFriendHandler() *FriendHandler {
	return &FriendHandler{}
}

func (h *FriendHandler) HandleGetFriends(c *fiber.Ctx) error {

	return nil
}
