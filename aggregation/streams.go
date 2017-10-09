package aggregation

import (
	"net/http"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/api/streams"
)

// Streams aggregation contains requirement to handle the process
type Streams struct {
	Context
}

func (s Streams) processStreams() {
	s.Loggger.Log("Aggregate on Streams Summary...")
	for _, credential := range s.Context.Credentials {
		twitchRequest := core.NewUserAccessTokenRequest(s.OAuth2, credential.TokenResponse)
		twitchRequest.Logger = s.Loggger
		twitchRequest.Method = http.MethodGet
		lastChannelEntry := s.Context.DB.GetLastUpdatedChannelSummary(credential.ChannelName)
		streamsManager := streams.Streams{Request: twitchRequest}
		streamsManager.GetStream(lastChannelEntry.IDTwitch, streams.Live)
	}
}

// Process aggregation
func (s Streams) Process() {
	s.Loggger.Log("Start Streams aggregation...")
	s.processStreams()
}
