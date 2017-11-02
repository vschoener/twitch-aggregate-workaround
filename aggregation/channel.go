package aggregation

import (
	"fmt"

	"github.com/wonderstream/twitch/core"
	coreModel "github.com/wonderstream/twitch/core/model"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// Channel aggregation contains requirement to handle the process
type Channel struct {
	Aggregator
	a *Aggregation
	repository.ChannelRepository
	repository.ChannelVideoRepository
	channelService      service.ChannelService
	subscriptionService service.SubscriptionService
	videoService        service.VideoService
}

// Initialize channel aggregator
func (c *Channel) Initialize(a *Aggregation) {
	c.a = a

	commonRepository := repository.NewRepository(a.Database, a.Logger)
	c.ChannelRepository = repository.ChannelRepository{
		Repository: commonRepository,
	}
	c.ChannelVideoRepository = repository.ChannelVideoRepository{
		Repository: commonRepository,
	}

	c.channelService = service.NewChannelService()
	c.subscriptionService = service.NewSubscriptionService()
	c.videoService = service.NewVideoService()
}

// Process channel aggregator
func (c Channel) Process(u model.User, isAuthenticated bool, token core.TokenResponse) {
	twitchRequest := core.NewAccessTokenRequest(c.a.OAuth2, token)
	twitchRequest.Logger = c.a.Logger.Share()
	twitchRequest.Logger.SetPrefix("LIBRARY")

	c.a.Logger.Log(fmt.Sprintf("Start Channel aggregation on %s #%d", u.Name, u.UserID))
	c.updateChannelSummary(u, isAuthenticated, twitchRequest)
	c.updateSubscriptionSummary(twitchRequest)
	c.GetVideosStream(twitchRequest, u.UserID)
}

// End channel aggregator
func (c Channel) End() {

}

func (c Channel) updateChannelSummary(u model.User, isAuthenticated bool, r *core.Request) {
	c.a.Logger.Log("Get Info")

	var err error
	var channel coreModel.Channel

	if isAuthenticated {
		channel, err = c.channelService.GetInfo(r)
	} else {
		channel, err = c.channelService.GetInfoByID(u.UserID, r)
	}

	if err != nil {
		return
	}
	c.StoreChannel(transformer.TransformCoreChannelToStorageChannel(channel))
}

func (c Channel) updateSubscriptionSummary(r *core.Request) {
	c.a.Logger.Log("Get Subscription")

	// subscriptionRepository := repository.SubscriptionRepository{
	// 	Repository: repository.NewRepository(c.DB, c.Loggger),
	// }

	//s := c.subscriptionService.GetSubscription(cr.ChannelID, userAccessTokenRequest)
	// TODO: Store using repository
}

// GetVideosStream retrieves last stream video information
func (c Channel) GetVideosStream(r *core.Request, userID int64) {
	c.a.Logger.Log("Get Video Stream")
	videos := c.videoService.GetVideosFromID(userID, c.a.twPublicRequest, 100)
	for _, video := range videos {
		c.RegisterVideoToChannel(userID, transformer.TransformCoreVideoToStorageVideo(video))
	}
}
