package repositories

import (
	"errors"
	"split-it/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserBySubID(subID string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "sub_id = ?", subID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUsersByName(name string) ([]models.User, error) {
	var users []models.User
	if err := r.db.Where("name LIKE ?", "%"+name+"%").Limit(50).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
