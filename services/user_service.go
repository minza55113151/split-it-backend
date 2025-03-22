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

func (s *UserService) CreateUser(uid string) (*models.User, error) {
	user := &models.User{
		UID: uid,
	}

	res, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) GetUserByUID(uid string) (*models.User, error) {
	user, err := s.repo.GetUserByUID(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(user *models.User) error {
	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id string) error {
	if err := s.repo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
