package model

import (
	"time"
)

// StreamType define an enum to use all type available
type StreamType int

const (
	// Live is the current stream
	Live StreamType = iota
	// Playlist is offline streams of VODs (Video on Demand) that appear live.
	Playlist
	// All returns everything
	All
)

var streamTypes = []string{
	"live",
	"playlist",
	"all",
}

// String returns the readable Enum
func (st StreamType) String() string {
	return streamTypes[st]
}

// Stream represent the JSON response from Twitch
type Stream struct {
	ID          int64             `json:"_id"`
	Game        string            `json:"game"`
	Viewers     int64             `json:"viewers"`
	VideoHeight int8              `json:"video_height"`
	AverageFPS  int8              `json:"average_fps"`
	Delay       int8              `json:"delay"`
	CreatedAt   time.Time         `json:"created_at"`
	IsPlaylist  bool              `json:"is_playlist"`
	Preview     map[string]string `json:"preview"`
	Channel     `json:"channel"`
}
