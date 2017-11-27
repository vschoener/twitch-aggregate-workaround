package model

import (
	"time"
)

const (
	// ChannelVideoTable database table
	ChannelVideoTable = "channel_videos"
)

// ChannelVideo mapping table
type ChannelVideo struct {
	ID              int64
	MetaDateAdd     time.Time
	ChannelID       int64
	Title           string
	Description     string
	DescriptionHTML string
	BroadcastID     int64
	BroadcastType   string
	Status          string
	TagList         string
	Views           int64
	URL             string
	Language        string
	CreatedAt       time.Time
	Viewable        string
	ViewableAt      string
	PublishedAt     time.Time
	VideoID         string
	RecordedAt      time.Time
	Game            string
	Length          int64
}
