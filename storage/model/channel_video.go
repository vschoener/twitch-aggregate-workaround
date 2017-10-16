package model

import (
	"time"
)

const (
	// ChannelVideoTable database table
	ChannelVideoTable = "channel_video"
)

// ChannelVideo mapping table
type ChannelVideo struct {
	ID              int64
	ChannelID       int64
	DateAdd         time.Time
	Title           string
	Description     string
	DescriptionHTML string
	BrodcastID      int64
	BrodcastType    string
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
