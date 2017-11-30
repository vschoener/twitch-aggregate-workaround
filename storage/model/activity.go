package model

import (
	"time"

	"github.com/wonderstream/twitch/storage"
)

const (
	// ActivityTable database table
	ActivityTable = "activity"

	// ActivityDatabase database name
	ActivityDatabase = storage.DBActivity
)

// Activity mapping table
type Activity struct {
	ID         int64
	DateTime   time.Time `gorm:"column:datetime"`
	Channel    string
	Username   string
	Type       string
	Content    string
	ContenHash string `gorm:"column:content_fnv1a_hash"`
}

// TableName set be singular
func (Activity) TableName() string {
	return ActivityTable
}
