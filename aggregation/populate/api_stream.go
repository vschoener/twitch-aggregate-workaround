package populate

import (
	"github.com/wonderstream/twitch/aggregation/service"
	"github.com/wonderstream/twitch/core"
	twModel "github.com/wonderstream/twitch/core/model"
	twService "github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage/model"
)

// Stream aggregation contains requirement to handle the process
type Stream struct {
	Loader        *service.Loader
	streamService twService.StreamService
}

// Initialize channel aggregator
func (s *Stream) Initialize(loader *service.Loader) {
	s.Loader = loader
	s.streamService = twService.NewStreamService()
}

func (s Stream) processStreams(userID int64, r *core.Request) {
	s.Loader.Logger.Log("Get Stream")
	s.streamService.GetStream(userID, twModel.Live, r)
}

// Process aggregation
func (s Stream) Process(user model.User, token core.TokenResponse) {
	twitchRequest := core.NewRequest(s.Loader.OAuth2, &token)
	twitchRequest.Logger = s.Loader.Logger.Share()
	twitchRequest.Logger.SetPrefix("LIBRARY")

	s.processStreams(user.ID, twitchRequest)
}

// End aggregate user
func (s Stream) End() {

}
