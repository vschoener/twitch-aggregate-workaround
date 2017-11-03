package model

import (
	"time"
)

const (
	// SummarizeTable database table
	SummarizeTable = "summarizes"
)

// Summarize mapping table
type Summarize struct {
	ID           int64
	ChannelID    int64
	ChannelName  string
	AVGCCV       int64 `gorm:"column:avg_ccv"`
	MaxCCV       int64
	AirTime      int64
	HoursWatched int64
	PrimaryGame  string
	Partner      bool
	Mature       bool
	Language     string
	Views        int64
	Followers    int64
	DateAdd      time.Time
}
