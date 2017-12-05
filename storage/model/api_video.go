package model

import (
	"time"
)

const (
	// VideoTable database table
	VideoTable = "api_videos"
)

// Video mapping table
type Video struct {
	MetaID          int64 `gorm:"primary_key:true"`
	MetaDateAdd     time.Time
	MetalChannelID  int64
	ID              string `gorm:"unique_index"`
	Title           string
	Description     string `gorm:"type:TEXT"`
	DescriptionHTML string `gorm:"type:TEXT"`
	BroadcastID     int64
	BroadcastType   string
	Status          string
	TagList         string `gorm:"type:TEXT"`
	Views           int64
	URL             string
	Language        string
	CreatedAt       time.Time
	Viewable        string
	ViewableAt      string
	PublishedAt     time.Time
	RecordedAt      time.Time
	Game            string
	Length          int64
}

// TableName set be singular
func (Video) TableName() string {
	return VideoTable
}
