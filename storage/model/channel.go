package model

import (
	"time"

	"github.com/wonderstream/twitch/core/model"
)

const (
	// ChannelTable database table
	ChannelTable = "channel"
)

// Channel mapping table
type Channel struct {
	ID      int64
	DateAdd time.Time
	model.Channel
}
