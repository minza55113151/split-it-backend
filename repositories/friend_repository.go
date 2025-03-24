package repositories

import (
	"errors"
	"split-it/models"

	"gorm.io/gorm"
)

type FriendRepository struct {
	db *gorm.DB
}

func NewFriendRepository(db *gorm.DB) *FriendRepository {
	return &FriendRepository{db: db}
}

func (r *FriendRepository) GetFriends(subID string) ([]models.FriendResponse, error) {
	var friends []models.FriendResponse

	err := r.db.Table("friends").
		Select("friends.*, users.name AS name, users.email AS email, users.image_url AS image_url").
		Where("friends.sub_id1 = ? OR friends.sub_id2 = ?", subID, subID).
		Joins("JOIN users ON (friends.sub_id1 = users.sub_id AND friends.sub_id2 = ?) OR (friends.sub_id2 = users.sub_id AND friends.sub_id1 = ?)", subID, subID).
		Where("users.sub_id != ?", subID).
		Find(&friends).Error

	return friends, err
}

func (r *FriendRepository) CreateFriend(subID1, subID2 string) error {
	subID1, subID2 = min(subID1, subID2), max(subID1, subID2)

	friend := &models.Friend{
		SubID1: subID1,
		SubID2: subID2,
		Status: "friend",
	}

	var count int64
	err := r.db.Table("users").Where("sub_id = ?", subID1).Or("sub_id = ?", subID2).Count(&count).Error
	if err != nil {
		return err
	}
	if count != 2 {
		return errors.New("one or both subIDs do not belong to valid users")
	}

	return r.db.Create(friend).Error
}

func (r *FriendRepository) DeleteFriend(subID1, subID2 string) error {
	subID1, subID2 = min(subID1, subID2), max(subID1, subID2)

	return r.db.Where("sub_id1 = ? AND sub_id2 = ?", subID1, subID2).Delete(&models.Friend{}).Error
}
