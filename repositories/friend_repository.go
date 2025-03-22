package repositories

import (
	"split-it/models"

	"gorm.io/gorm"
)

type FriendRepository struct {
	db *gorm.DB
}

func NewFriendRepository(db *gorm.DB) *FriendRepository {
	return &FriendRepository{db: db}
}

func (r *FriendRepository) GetFriends(subID string) ([]models.Friend, error) {
	var friends []models.Friend

	err := r.db.Where("sub_id1 = ? OR sub_id2 = ?", subID, subID).Find(&friends).Error
	// TODO: map
	return friends, err
}

func (r *FriendRepository) CreateFriend(subID1, subID2 string) error {
	subID1, subID2 = min(subID1, subID2), max(subID1, subID2)

	friend := &models.Friend{
		SubID1: subID1,
		SubID2: subID2,
		Status: "friend",
	}

	return r.db.Create(friend).Error
}

func (r *FriendRepository) DeleteFriend(subID1, subID2 string) error {
	subID1, subID2 = min(subID1, subID2), max(subID1, subID2)

	return r.db.Where("sub_id1 = ? AND sub_id2 = ?", subID1, subID2).Delete(&models.Friend{}).Error
}
