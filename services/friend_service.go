package services

import "split-it/repositories"

type FriendService struct {
	friendRepo *repositories.FriendRepository
}

func NewFriendService(friendRepo *repositories.FriendRepository) *FriendService {
	return &FriendService{
		friendRepo: friendRepo,
	}
}
