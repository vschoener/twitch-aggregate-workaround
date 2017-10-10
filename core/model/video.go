package model

import "time"

// Video is the model from Twitch API
type Video struct {
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
	VideoID         string    `json:"_id"`
	RecordedAt      time.Time `json:"recorded_at"`
	Game            string    `json:"game"`
	Length          int64     `json:"length"`
}
