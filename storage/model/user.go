package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	// UserTable database table
	UserTable = "users"
)

// User is the storage manager
type User struct {
	gorm.Model
	DisplayName string
	UserID      int64 `gorm:"unique_index"`
	Name        string
	Type        string
	Bio         string
	Logo        string
	DateAdd     time.Time
}
