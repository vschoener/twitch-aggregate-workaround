package aggregation

import (
	"github.com/wonderstream/twitch/core"
	coreModel "github.com/wonderstream/twitch/core/model"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage/model"
)

// Stream aggregation contains requirement to handle the process
type Stream struct {
	Aggregator
	a             *Aggregation
	streamService service.StreamService
}

// Initialize channel aggregator
func (s *Stream) Initialize(a *Aggregation) {
	s.a = a
	s.streamService = service.NewStreamService()
}

func (s Stream) processStreams(userID int64) {
	s.a.Logger.Log("Get Stream")
	s.streamService.GetStream(userID, coreModel.Live, s.a.twPublicRequest)
}

// Process aggregation
func (s Stream) Process(user model.User, isAuthenticated bool, token core.TokenResponse) {
	s.processStreams(user.UserID)
}

// End aggregate user
func (s Stream) End() {

}
