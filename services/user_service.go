package services

import (
	"split-it/models"
	"split-it/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(subID string) (*models.User, error) {
	user := &models.User{
		SubID: subID,
	}

	res, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) GetUserBySubID(subID string) (*models.User, error) {
	user, err := s.repo.GetUserBySubID(subID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
