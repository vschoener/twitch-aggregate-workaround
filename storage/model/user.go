package model

import (
	"time"
)

const (
	// UserTable database table
	UserTable = "user"
)

// User is the storage manager
type User struct {
	ID          int64
	DateAdd     time.Time
	DisplayName string
	UserID      int64
	Name        string
	Type        string
	Bio         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Logo        string
}
