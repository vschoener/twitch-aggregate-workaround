package aggregation

import (
	"net/http"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage/repository"
)

// Streams aggregation contains requirement to handle the process
type Streams struct {
	Context
	repository.ChannelRepository
}

func (s Streams) processStreams() {
	s.Loggger.Log("Aggregate on Streams Summary...")
	for _, credential := range s.Context.Credentials {
		twitchRequest := core.NewUserAccessTokenRequest(s.OAuth2, credential.TokenResponse)
		twitchRequest.Logger = s.Loggger
		twitchRequest.Method = http.MethodGet
		lastChannelEntry := s.GetLastUpdatedChannelSummary(credential.ChannelName)
		streamService := service.StreamService{Request: twitchRequest}
		streamService.GetStream(lastChannelEntry.IDTwitch, model.Live)
	}
}

// Process aggregation
func (s Streams) Process() {
	s.Loggger.Log("Start Streams aggregation...")
	s.processStreams()
}
