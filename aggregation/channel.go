package aggregation

import (
	"fmt"
	"log"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// Channel aggregation contains requirement to handle the process
type Channel struct {
	Context
	repository.ChannelRepository
	repository.ChannelVideoRepository
}

func (c Channel) updateChannelSummary() {
	c.Loggger.Log("Aggregate on Channel Summary...")

	for _, credential := range c.Context.Credentials {
		twitchRequest := core.NewUserAccessTokenRequest(c.OAuth2, credential.TokenResponse)
		twitchRequest.Logger = c.Loggger
		channelService := service.ChannelService{Request: twitchRequest}
		channel := channelService.GetInfo()

		c.StoreChannel(transformer.TransformCoreChannelToStorageChannel(channel))
	}
}

func (c Channel) updateSubscriptionSummary() {
	c.Loggger.Log("Aggregate on Subscription Summary...")

	// subscriptionRepository := repository.SubscriptionRepository{
	// 	Repository: repository.NewRepository(c.DB, c.Loggger),
	// }

	for _, credential := range c.Context.Credentials {
		twitchRequest := core.NewUserAccessTokenRequest(c.OAuth2, credential.TokenResponse)
		twitchRequest.Logger = c.Loggger
		subService := service.SubscriptionService{Request: twitchRequest}
		channel := c.GetLastUpdatedChannelSummary(credential.ChannelName)
		s := subService.GetSubscription(channel.IDTwitch)

		log.Println(fmt.Sprintf("Not used: %#v", s))
	}
}

// GetVideosStream retrieves last stream video information
func (c Channel) GetVideosStream() {
	channels := c.GetChannels()
	for _, channel := range channels {
		videoService := service.VideoService{
			Request: c.Request,
		}

		videos := videoService.GetVideosFromID(channel.IDTwitch, 100)
		for _, video := range videos {
			c.RegisterVideoToChannel(channel.IDTwitch, transformer.TransformCoreVideoToStorageVideo(video))
		}
	}
}

// Process aggregation
func (c Channel) Process() {
	c.Loggger.Log("Start Channel aggregation...")
	c.updateChannelSummary()
	//c.updateSubscriptionSummary()
	c.GetVideosStream()
}
