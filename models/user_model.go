package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	UID       string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
