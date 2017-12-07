package precompute

import (
	"time"

	"github.com/wonderstream/twitch/aggregation/service"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
	sModel "github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
)

// Channel aggregation contains requirement to handle the process
type Channel struct {
	Loader              *service.Loader
	sRepo               repository.PrecomputeChannelRepository
	cRepo               repository.ChannelRepository
	cvRepo              repository.VideoRepository
	activityStorageRepo repository.ActivityStorageRepository
	caService           service.ChannelService
	PrecomputedChannel  sModel.PrecomputedChannel
	QueryFilter         storage.QueryFilter
}

// Initialize channel aggregator
func (s *Channel) Initialize(loader *service.Loader) {
	s.Loader = loader
	s.sRepo = repository.NewPrecomputeChannelRepository(s.Loader.DatabaseManager.Get(storage.DBAggregation), s.Loader.Logger)
	s.cRepo = repository.NewChannelRepository(s.Loader.DatabaseManager.Get(storage.DBAggregation), s.Loader.Logger)
	s.cvRepo = repository.NewChannelVideoRepository(s.Loader.DatabaseManager.Get(storage.DBAggregation), s.Loader.Logger)
	s.activityStorageRepo = repository.NewActivityStorageRepository(s.Loader.DatabaseManager.Get(storage.DBActivity), s.Loader.Logger)
	s.activityStorageRepo.Database.Gorm.LogMode(true)
}

// Process channel aggregator
func (s Channel) Process(u sModel.User, token core.TokenResponse) {

	// Example to use your own date
	// dateStart := time.Date(2017, time.October, 20, 0, 0, 0, 0, time.UTC)
	// dateEnd := time.Date(2017, time.October, 21, 0, 0, 0, 0, time.UTC)

	// The DateField needs to be set individualy inside the right compute
	// method call (cause it mays change depending of the twitch record)

	if channel, found := s.cRepo.GetLastRecorded(u.Name); found {
		// aggregate api computing
		s.PrecomputedChannel.DateAdd = time.Now()
		s.PrecomputedChannel.Followers = channel.Followers
		s.PrecomputedChannel.Views = channel.Views
		s.PrecomputedChannel.ChannelID = channel.ID
		s.PrecomputedChannel.ChannelName = channel.Name
		s.PrecomputedChannel.Partner = channel.Partner
		s.PrecomputedChannel.Mature = channel.Mature
		s.PrecomputedChannel.Language = channel.Language
		s.PrecomputedChannel.PrimaryGame = s.caService.GetPrimaryGame(channel.ID, s.QueryFilter, s.cvRepo).Name
		s.PrecomputedChannel.AirTime = s.cvRepo.GetAirTime(channel.ID, s.QueryFilter)

		// activity computing
		s.PrecomputedChannel.UniqueViewers = s.activityStorageRepo.GetUniqueViewers(channel.Name, s.QueryFilter)
		s.PrecomputedChannel.SecondsWatched = s.activityStorageRepo.GetWatchedTime(channel.Name, s.QueryFilter)
		ccvInfo := s.activityStorageRepo.GetCCVInformation(channel.Name, s.QueryFilter)
		s.PrecomputedChannel.MaxCCV = ccvInfo.MaxCCV
		s.PrecomputedChannel.AVGCCV = ccvInfo.AvgCCV
		err := s.sRepo.Add(s.PrecomputedChannel)

		if err != nil {
			s.Loader.Logger.LogInterface(err)
		}
	}
}

// End channel aggregator
func (s Channel) End() {

}
