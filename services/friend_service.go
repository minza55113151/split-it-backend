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

func (s *FriendService) GetFriends(subID string) ([]models.FriendResponse, error) {
	friends, err := s.friendRepo.GetFriends(subID)
	if err != nil {
		return nil, err
	}

	for i := range friends {
		if friends[i].SubID1 == subID {
			friends[i].SubID = friends[i].SubID2
		} else {
			friends[i].SubID = friends[i].SubID1
		}
	}

	return friends, nil
}

func (s *FriendService) CreateFriend(subID1, subID2 string) error {
	return s.friendRepo.CreateFriend(subID1, subID2)
}

func (s *FriendService) DeleteFriend(subID1, subID2 string) error {
	return s.friendRepo.DeleteFriend(subID1, subID2)
}
