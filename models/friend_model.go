package models

import (
	"time"
)

type Friend struct {
	ID        uint      `gorm:"primaryKey"`
	SubID1    string    `gorm:"uniqueIndex:idx_friend_unique"`
	SubID2    string    `gorm:"uniqueIndex:idx_friend_unique"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
