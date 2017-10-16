package model

import (
	"time"
)

const (
	// ChannelTable database table
	ChannelTable = "channel"
)

// Channel mapping table
type Channel struct {
	ID                   int64
	DateAdd              time.Time
	Mature               bool   `json:"mature"`
	Status               string `json:"status"`
	BroadcasterLanguage  string `json:"broadcaster_language"`
	DisplayName          string `json:"display_name"`
	Game                 string `json:"game"`
	Language             string `json:"language"`
	IDTwitch             int64  `json:"_id,string"`
	Name                 string `json:"name"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	Partner              bool   `json:"partner"`
	Logo                 string `json:"logo"`
	VideoBanner          string `json:"video_banner"`
	ProfileBanner        string `json:"profile_banner"`
	ProfileBannerBGColor string `json:"profile_banner_background_color"`
	URL                  string `json:"url"`
	Views                int64  `json:"views"`
	Followers            int64  `json:"followers"`
	BroadcasterType      string `json:"broadcaster_type"`
	StreamKey            string `json:"stream_key"`
	Email                string `json:"email"`
}
