package services

import (
	"split-it/models"
	"split-it/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateUser(subID string) (*models.User, error) {
	user := &models.User{
		SubID: subID,
	}

	res, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) GetUserBySubID(subID string) (*models.User, error) {
	user, err := s.userRepo.GetUserBySubID(subID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
