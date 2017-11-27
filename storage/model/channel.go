package model

import (
	"time"
)

const (
	// ChannelTable database table
	ChannelTable = "channels"
)

// Channel mapping table
type Channel struct {
	ID                   int64
	MetaDateAdd          time.Time
	Mature               bool
	Status               string
	BroadcasterLanguage  string
	DisplayName          string
	Game                 string
	Language             string
	ChannelID            int64
	Name                 string
	CreatedAt            string
	UpdatedAt            string
	Partner              bool
	Logo                 string
	VideoBanner          string
	ProfileBanner        string
	ProfileBannerBGColor string
	URL                  string
	Views                int64
	Followers            int64
	BroadcasterType      string
	StreamKey            string
	Email                string
}
