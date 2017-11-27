package aggregation

import (
	"time"

	"github.com/wonderstream/twitch/aggregation/service"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
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
	caService service.ChannelService
	Summarize sModel.Summarize
}

// Initialize channel aggregator
func (s *Summarize) Initialize(a *Aggregation) {
	s.a = a
	s.sRepo = repository.NewSummarizeRepository(a.DM.Get(storage.DBAggregation), a.Logger)
	s.cRepo = repository.NewChannelRepository(a.DM.Get(storage.DBAggregation), a.Logger)
	s.cvRepo = repository.NewChannelVideoRepository(a.DM.Get(storage.DBAggregation), a.Logger)
}

// Process channel aggregator
func (s Summarize) Process(u sModel.User, isAuthenticated bool, token core.TokenResponse) {
	if nil == s.a {
		panic("Aggregation missing, did you forget to set it ?")
	}

	dateStart := time.Date(2017, time.October, 20, 0, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2017, time.October, 21, 0, 0, 0, 0, time.UTC)
	queryFilter := storage.QueryFilter{
		DateStart: &dateStart,
		DateEnd:   &dateEnd,
	}
	if channel, found := s.cRepo.GetLastRecorded(u.Name); found {
		s.Summarize.Followers = channel.Followers
		s.Summarize.Views = channel.Views
		s.Summarize.ChannelID = channel.ChannelID
		s.Summarize.ChannelName = channel.Name
		s.Summarize.Partner = channel.Partner
		s.Summarize.Mature = channel.Mature
		s.Summarize.Language = channel.Language
		s.Summarize.AirTime = s.cvRepo.GetAirTime(channel.ChannelID, queryFilter)
		s.Summarize.PrimaryGame = s.caService.GetPrimaryGame(channel.ChannelID, queryFilter, s.cvRepo).Name

		err := s.a.DM.Get(storage.DBAggregation).Gorm.Create(&s.Summarize).Error

		if err != nil {
			s.a.Logger.LogInterface(err)
		}
	}
}

// End channel aggregator
func (s Summarize) End() {

}
