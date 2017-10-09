package core

import (
	"fmt"
	"time"
)

// Channel to manage anything about it
type Channel struct {
	*Request
}

const (
	// ChannelURI used when building Request
	ChannelURI = "/channel"

	// ChannelsURI used to build channels request
	ChannelsURI = "/channels"
)

// ChannelSummary contains channel information
type ChannelSummary struct {
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

// ChannelVideo contains the video information of this channel
type ChannelVideo struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	DescriptionHTML string    `json:"description_html"`
	BrodcastID      int64     `json:"broadcast_id"`
	BrodcastType    string    `json:"broadcast_type"`
	Status          string    `json:"status"`
	TagList         string    `json:"tag_list"`
	Views           int64     `json:"views"`
	URL             string    `json:"url"`
	Language        string    `json:"language"`
	CreatedAt       time.Time `json:"created_at"`
	Viewable        string    `json:"viewable"`
	ViewableAt      string    `json:"viewable_at"`
	PublishedAt     time.Time `json:"published_at"`
	ID              string    `json:"_id"`
	RecordedAt      time.Time `json:"recorded_at"`
	Game            string    `json:"game"`
	Length          int64     `json:"length"`
}

// NewChannel constructor
func NewChannel(r *Request) *Channel {
	return &Channel{
		Request: r,
	}
}

// RequestSummary retrieve information
func (c Channel) RequestSummary() ChannelSummary {

	channelSummary := ChannelSummary{}
	c.SendRequest(ChannelURI, &channelSummary)

	return channelSummary
}

// GetVideosFromIDResult contains result from request of GetVideosFromID
type GetVideosFromIDResult struct {
	Total  int16 `json:"_total"`
	Videos []ChannelVideo
}

// GetVideosFromID returns the videos list of the channel ID
func (c Channel) GetVideosFromID(channelID int64, total int) []ChannelVideo {
	result := GetVideosFromIDResult{}

	c.SendRequest(fmt.Sprintf("%s/%d/videos?limit=%d", ChannelsURI, channelID, total), &result)
	return result.Videos
}

// GetSubscriptionSummary return the subscription summary of the channel ID
func (c Channel) GetSubscriptionSummary(channelID int64) SubscriptionSummary {

	subscriptionSummary := SubscriptionSummary{}

	url := fmt.Sprintf("%s/%d/subscriptions", ChannelsURI, channelID)
	c.SendRequest(url, &subscriptionSummary)
	return subscriptionSummary
}
