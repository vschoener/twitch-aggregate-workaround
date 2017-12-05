package model

import "time"

const (
	// UserTable database table
	UserTable = "api_users"
)

// User is the storage manager
type User struct {
	MetaID         uint `gorm:"primary_key"`
	MetaDateAdd    time.Time
	MetaDateUpdate time.Time
	ID             int64 `gorm:"unique_index"`
	DisplayName    string
	Name           string
	Type           string
	Bio            string `gorm:"type:TEXT"`
	Logo           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// TableName set be singular
func (User) TableName() string {
	return UserTable
}
