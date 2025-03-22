package models

import (
	"time"
)

type Friend struct {
	ID        uint `gorm:"primaryKey"`
	UserID1   uint
	UserID2   uint
	Status    string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
