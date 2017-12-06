package model

import (
	"time"
)

const (
	// PrecomputedChannelTable database table
	PrecomputedChannelTable = "precomputed_channels"
)

// PrecomputedChannel mapping table
type PrecomputedChannel struct {
	ID             int64
	DateAdd        time.Time
	ChannelID      int64 `gorm:"ForeignKey:Videos"`
	ChannelName    string
	AVGCCV         int64 `gorm:"column:avg_ccv"`
	MaxCCV         int64
	AirTime        int64
	UniqueViewers  int64
	SecondsWatched int64
	PrimaryGame    string
	Partner        bool
	Mature         bool
	Language       string
	Views          int64
	Followers      int64
}

// TableName set be singular
func (PrecomputedChannel) TableName() string {
	return PrecomputedChannelTable
}
