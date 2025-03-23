package models

import "time"

// Expense represents an expense record in the system
type Expense struct {
	ID          uint    `gorm:"primaryKey"`
	PayerSubID  string  `gorm:"not null"`
	DebtorSubID string  `gorm:"not null"`
	Title       string  `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Icon        string  `gorm:"not null"`
	Currency    string  `gorm:"not null"`
	Note        string  `gorm:"not null"`
	Status      string  `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
