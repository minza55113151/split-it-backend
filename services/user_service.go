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

func (s *UserService) GetUserBySubID(subID string) (*models.User, error) {
	user, err := s.userRepo.GetUserBySubID(subID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUsersByName(name string) ([]models.User, error) {
	users, err := s.userRepo.GetUsersByName(name)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) CreateUser(subID string, name string, email string, imageURL string) (*models.User, error) {
	user := &models.User{
		SubID:    subID,
		Name:     name,
		Email:    email,
		ImageURL: imageURL,
	}

	res, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	oldUser, err := s.userRepo.GetUserBySubID(user.SubID)
	if err != nil {
		return nil, err
	}

	oldUser.Name = user.Name

	res, err := s.userRepo.UpdateUser(oldUser)
	if err != nil {
		return nil, err
	}

	return res, nil
}
