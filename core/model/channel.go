package model

import "time"

// Channel is the model from Twitch API
type Channel struct {
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	DisplayName                  string    `json:"display_name"`
	Game                         string    `json:"game"`
	Language                     string    `json:"language"`
	ID                           int64     `json:"_id,string"`
	Name                         string    `json:"name"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	Partner                      bool      `json:"partner"`
	Logo                         string    `json:"logo"`
	VideoBanner                  string    `json:"video_banner"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	URL                          string    `json:"url"`
	Views                        int64     `json:"views"`
	Followers                    int64     `json:"followers"`
	BroadcasterType              string    `json:"broadcaster_type"`
	StreamKey                    string    `json:"stream_key"`
	Email                        string    `json:"email"`
	Description                  string    `json:"description"`
}
