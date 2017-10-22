package aggregation

import (
	"fmt"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// Channel aggregation contains requirement to handle the process
type Channel struct {
	Aggregator
	a Aggregation
	repository.ChannelRepository
	repository.ChannelVideoRepository
	channelService      service.ChannelService
	subscriptionService service.SubscriptionService
	videoService        service.VideoService
}

// Initialize channel aggregator
func (c *Channel) Initialize(a Aggregation) {
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
func (c Channel) Process(cr model.Credential) {
	token := transformer.TransformStorageCredentialToCoreTokenResponse(cr)
	twitchRequest := core.NewUserAccessTokenRequest(c.a.OAuth2, token)
	twitchRequest.Logger = c.a.Logger.Share()
	twitchRequest.Logger.SetPrefix("LIBRARY")

	c.a.Logger.Log(fmt.Sprintf("Start Channel aggregation on %s #%d", cr.ChannelName, cr.ChannelID))
	c.updateChannelSummary(twitchRequest)
	c.updateSubscriptionSummary(cr, twitchRequest)
	c.GetVideosStream(cr)
}

// End channel aggregator
func (c Channel) End() {

}

func (c Channel) updateChannelSummary(userAccessTokenRequest *core.Request) {
	c.a.Logger.Log("Get Info")

	channel := c.channelService.GetInfo(userAccessTokenRequest)
	c.StoreChannel(transformer.TransformCoreChannelToStorageChannel(channel))
}

func (c Channel) updateSubscriptionSummary(cr model.Credential, userAccessTokenRequest *core.Request) {
	c.a.Logger.Log("Get Subscription")

	// subscriptionRepository := repository.SubscriptionRepository{
	// 	Repository: repository.NewRepository(c.DB, c.Loggger),
	// }

	//s := c.subscriptionService.GetSubscription(cr.ChannelID, userAccessTokenRequest)
	// TODO: Store using repository
}

// GetVideosStream retrieves last stream video information
func (c Channel) GetVideosStream(cr model.Credential) {
	c.a.Logger.Log("Get Video Stream")
	videos := c.videoService.GetVideosFromID(cr.ChannelID, c.a.twPublicRequest, 100)
	for _, video := range videos {
		c.RegisterVideoToChannel(cr.ChannelID, transformer.TransformCoreVideoToStorageVideo(video))
	}
}
