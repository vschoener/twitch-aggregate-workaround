package model

import "time"

const (
	// UserTable database table
	UserTable = "users"
)

// User is the storage manager
type User struct {
	ID             uint `gorm:"primary_key"`
	MetaDateAdd    time.Time
	MetaDateUpdate time.Time
	DisplayName    string
	UserID         int64 `gorm:"unique_index"`
	Name           string
	Type           string
	Bio            string
	Logo           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
