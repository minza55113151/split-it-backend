package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	SubID     string    `gorm:"uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
