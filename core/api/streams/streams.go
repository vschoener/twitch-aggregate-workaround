package streams

import (
	"fmt"
	"log"
	"time"

	"github.com/wonderstream/twitch/core"
)

// Streams manager
type Streams struct {
	*core.Request
}

const (
	// StreamURI is the end point to get stream
	StreamURI = "/streams"
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
	ID                  int64             `json:"_id"`
	Game                string            `json:"game"`
	Viewers             int64             `json:"viewers"`
	VideoHeight         int8              `json:"video_height"`
	AverageFPS          int8              `json:"average_fps"`
	Delay               int8              `json:"delay"`
	CreatedAt           time.Time         `json:"created_at"`
	IsPlaylist          bool              `json:"is_playlist"`
	Preview             map[string]string `json:"preview"`
	core.ChannelSummary `json:"channel"`
}

// GetStream Gets stream information (the stream object) for a specified user.
// Authentication none
func (s Streams) GetStream(channelID int64, streamType StreamType) {
	stream := struct {
		stream Stream
	}{}
	s.Request.SendRequest(fmt.Sprintf("%s/%d", StreamURI, channelID), stream)

	log.Println(stream)
}

// GetStreams Gets a list of live streams.
// Authentication none
func (s Streams) GetStreams() {

}

// GetStreamsSummary Gets a summary of live streams.
// Authentication none
func (s Streams) GetStreamsSummary() {

}

// GetFeaturedStreams Gets a list of all featured live streams..
// Authentication none
func (s Streams) GetFeaturedStreams() {

}

// GetFollowedStreams Gets a list of all featured live streams..
// Authentication Required scope: user_read
func (s Streams) GetFollowedStreams() {

}
