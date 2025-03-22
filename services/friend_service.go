package services

import (
	"split-it/models"
	"split-it/repositories"
)

type FriendService struct {
	friendRepo *repositories.FriendRepository
}

func NewFriendService(friendRepo *repositories.FriendRepository) *FriendService {
	return &FriendService{
		friendRepo: friendRepo,
	}
}

func (s *FriendService) GetFriends(subID string) ([]models.Friend, error) {
	return s.friendRepo.GetFriends(subID)
}

func (s *FriendService) CreateFriend(subID1, subID2 string) error {
	return s.friendRepo.CreateFriend(subID1, subID2)
}

func (s *FriendService) DeleteFriend(subID1, subID2 string) error {
	return s.friendRepo.DeleteFriend(subID1, subID2)
}
