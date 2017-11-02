package aggregation

import (
	"github.com/wonderstream/twitch/core"
	sModel "github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
)

// Summarize aggregation contains requirement to handle the process
type Summarize struct {
	Aggregator
	a         *Aggregation
	sRepo     repository.SummarizeRepository
	cRepo     repository.ChannelRepository
	cvRepo    repository.ChannelVideoRepository
	Summarize sModel.Summarize
}

// Initialize channel aggregator
func (s *Summarize) Initialize(a *Aggregation) {
	s.a = a
	s.sRepo = repository.NewSummarizeRepository(a.Database, a.Logger)
	s.cRepo = repository.NewChannelRepository(a.Database, a.Logger)
	s.cvRepo = repository.NewChannelVideoRepository(a.Database, a.Logger)
}

// Process channel aggregator
func (s Summarize) Process(u sModel.User, isAuthenticated bool, token core.TokenResponse) {
	if nil == s.a {
		panic("Aggregation missing, did you forget to set it ?")
	}

	if channel, found := s.cRepo.GetLastRecorded(u.Name); found {
		s.Summarize.Followers = channel.Followers
		s.Summarize.Views = channel.Views
		s.Summarize.ChannelID = channel.ChannelID
		s.Summarize.ChannelName = channel.Name
		s.Summarize.ChannelName = channel.Name
		s.Summarize.Partner = channel.Partner
		s.Summarize.Mature = channel.Mature
		s.Summarize.Language = channel.Language
		s.Summarize.AirTime = s.cvRepo.GetAirTime(channel.ChannelID, nil, nil)

		err := s.a.Database.Gorm.Create(&s.Summarize).Error

		if err != nil {
			s.a.Logger.LogInterface(err)
		}
	}
}

// End channel aggregator
func (s Summarize) End() {

}
