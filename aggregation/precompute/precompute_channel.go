package aggregation

import (
	"time"

	"github.com/wonderstream/twitch/aggregation/service"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
	sModel "github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
)

// PrecomputeChannel aggregation contains requirement to handle the process
type PrecomputeChannel struct {
	Loader              *service.Loader
	sRepo               repository.PrecomputeChannelRepository
	cRepo               repository.ChannelRepository
	cvRepo              repository.VideoRepository
	activityStorageRepo repository.ActivityStorageRepository
	caService           service.ChannelService
	PrecomputedChannel  sModel.PrecomputedChannel
}

// Initialize channel aggregator
func (s *PrecomputeChannel) Initialize(loader *service.Loader) {
	s.Loader = loader
	s.sRepo = repository.NewPrecomputeChannelRepository(s.Loader.DatabaseManager.Get(storage.DBAggregation), s.Loader.Logger)
	s.cRepo = repository.NewChannelRepository(s.Loader.DatabaseManager.Get(storage.DBAggregation), s.Loader.Logger)
	s.cvRepo = repository.NewChannelVideoRepository(s.Loader.DatabaseManager.Get(storage.DBAggregation), s.Loader.Logger)
	s.activityStorageRepo = repository.NewActivityStorageRepository(s.Loader.DatabaseManager.Get(storage.DBActivity), s.Loader.Logger)
}

// Process channel aggregator
func (s PrecomputeChannel) Process(u sModel.User, token core.TokenResponse) {

	dateStart := time.Date(2017, time.October, 20, 0, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2017, time.October, 21, 0, 0, 0, 0, time.UTC)
	queryFilter := storage.QueryFilter{
		DateStart: &dateStart,
		DateEnd:   &dateEnd,
	}
	if channel, found := s.cRepo.GetLastRecorded(u.Name); found {
		s.PrecomputedChannel.MetaDateAdd = time.Now()
		s.PrecomputedChannel.Followers = channel.Followers
		s.PrecomputedChannel.Views = channel.Views
		s.PrecomputedChannel.ChannelID = channel.ID
		s.PrecomputedChannel.ChannelName = channel.Name
		s.PrecomputedChannel.Partner = channel.Partner
		s.PrecomputedChannel.Mature = channel.Mature
		s.PrecomputedChannel.Language = channel.Language
		s.PrecomputedChannel.AirTime = s.cvRepo.GetAirTime(channel.ID, queryFilter)
		s.PrecomputedChannel.PrimaryGame = s.caService.GetPrimaryGame(channel.ID, queryFilter, s.cvRepo).Name
		s.PrecomputedChannel.SecondsWatched = s.activityStorageRepo.GetWatchedTime(channel.Name, queryFilter)

		err := s.sRepo.Add(s.PrecomputedChannel)

		if err != nil {
			s.Loader.Logger.LogInterface(err)
		}
	}
}

// End channel aggregator
func (s PrecomputeChannel) End() {

}
