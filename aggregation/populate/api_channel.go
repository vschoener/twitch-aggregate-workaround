package populate

import (
	"fmt"
	"log"

	"github.com/wonderstream/twitch/aggregation/service"
	twCore "github.com/wonderstream/twitch/core"
	twModel "github.com/wonderstream/twitch/core/model"
	twService "github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// Channel aggregation contains requirement to handle the process
type Channel struct {
	Loader *service.Loader
	repository.ChannelRepository
	repository.VideoRepository
	channelService      twService.ChannelService
	subscriptionService twService.SubscriptionService
	videoService        twService.VideoService
}

// Initialize channel aggregator
func (c *Channel) Initialize(l *service.Loader) {
	c.Loader = l
	commonRepository := repository.NewRepository(l.DatabaseManager.Get(storage.DBAggregation), l.Logger)
	c.ChannelRepository = repository.ChannelRepository{
		Repository: commonRepository,
	}
	c.VideoRepository = repository.VideoRepository{
		Repository: commonRepository,
	}

	c.channelService = twService.NewChannelService()
	c.subscriptionService = twService.NewSubscriptionService()
	c.videoService = twService.NewVideoService()
}

// Process channel aggregator
func (c Channel) Process(u model.User, token twCore.TokenResponse) {
	twitchRequest := twCore.NewRequest(c.Loader.OAuth2, &token)
	twitchRequest.Logger = c.Loader.Logger.Share()
	twitchRequest.Logger.SetPrefix("LIBRARY")

	c.Loader.Logger.Log(fmt.Sprintf("Start Channel aggregation on %s #%d", u.Name, u.ID))
	c.updateChannelSummary(u, token.IsAuthenticated(), twitchRequest)
	//c.updateSubscriptionSummary(twitchRequest)
	c.GetVideosStream(twitchRequest, u.ID)
}

// End channel aggregator
func (c Channel) End() {

}

func (c Channel) updateChannelSummary(u model.User, isAuthenticated bool, r *twCore.Request) {
	c.Loader.Logger.Log("Get Info")

	var err error
	var channel twModel.Channel

	if isAuthenticated {
		channel, err = c.channelService.GetInfo(r)
	} else {
		channel, err = c.channelService.GetInfoByID(u.ID, r)
	}

	if err != nil {
		return
	}
	c.StoreChannel(transformer.TransformCoreChannelToStorageChannel(channel))
}

func (c Channel) updateSubscriptionSummary(r *twCore.Request) {
	c.Loader.Logger.Log("Get Subscription")

	// subscriptionRepository := repository.SubscriptionRepository{
	// 	Repository: repository.NewRepository(c.DB, c.Loggger),
	// }

	//s := c.subscriptionService.GetSubscription(cr.ChannelID, userAccessTokenRequest)
	// TODO: Store using repository
}

// GetVideosStream retrieves last stream video information
func (c Channel) GetVideosStream(r *twCore.Request, userID int64) {
	c.Loader.Logger.Log("Get Video Stream")
	videos := c.videoService.GetVideosFromID(userID, r, 100)
	sVideos := []model.Video{}
	for _, video := range videos {
		log.Println("Link videos", video)
		sVideos = append(sVideos, transformer.TransformCoreVideoToStorageVideo(video))
	}
	//c.VideoRepository.(userID)
}
