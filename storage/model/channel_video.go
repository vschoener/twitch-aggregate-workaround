package model

import (
	"time"

	"github.com/wonderstream/twitch/core/model"
)

const (
	// ChannelVideoTable database table
	ChannelVideoTable = "channel_video"
)

// ChannelVideo mapping table
type ChannelVideo struct {
	ID        int64
	ChannelID int64
	DateAdd   time.Time
	model.Video
}
