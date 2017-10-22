package aggregation

import (
	coreModel "github.com/wonderstream/twitch/core/model"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage/model"
)

// Stream aggregation contains requirement to handle the process
type Stream struct {
	Aggregator
	a             Aggregation
	streamService service.StreamService
}

// Initialize channel aggregator
func (s *Stream) Initialize(a Aggregation) {
	s.a = a
	s.streamService = service.NewStreamService()
}

func (s Stream) processStreams(cr model.Credential) {
	s.a.Logger.Log("Get Stream")
	s.streamService.GetStream(cr.ChannelID, coreModel.Live, s.a.twPublicRequest)
}

// Process aggregation
func (s Stream) Process(cr model.Credential) {
	s.processStreams(cr)
}

// End aggregate user
func (s Stream) End() {

}
