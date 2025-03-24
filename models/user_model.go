package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	SubID     string `gorm:"uniqueIndex"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex"`
	ImageURL  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
