package service

import (
	"fmt"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
)

// StreamService handles processes for the stream
type StreamService struct {
	*core.Request
}

// GetStream Gets stream information (the stream object) for a specified user.
// Authentication none
func (s StreamService) GetStream(channelID int64, streamType model.StreamType) model.Stream {
	stream := struct {
		stream model.Stream
	}{}
	s.Request.SendRequest(fmt.Sprintf("%s/%d", core.StreamsURI, channelID), stream)

	return stream.stream
}

// GetStreams Gets a list of live streams.
// Authentication none
func (s StreamService) GetStreams() {

}

// GetStreamsSummary Gets a summary of live streams.
// Authentication none
func (s StreamService) GetStreamsSummary() {

}

// GetFeaturedStreams Gets a list of all featured live streams..
// Authentication none
func (s StreamService) GetFeaturedStreams() {

}

// GetFollowedStreams Gets a list of all featured live streams..
// Authentication Required scope: user_read
func (s StreamService) GetFollowedStreams() {

}
