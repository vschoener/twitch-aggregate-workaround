package aggregation

import (
	"fmt"
	"log"

	"net/http"

	"github.com/wonderstream/twitch/core"
)

// Channel aggregation contains requirement to handle the process
type Channel struct {
	Context
}

func (c Channel) updateChannelSummary() {
	c.Loggger.Log("Aggregate on Channel Summary...")
	for _, credential := range c.Context.Credentials {
		twitchRequest := core.NewUserAccessTokenRequest(c.OAuth2, credential.TokenResponse)
		twitchRequest.Logger = c.Loggger
		twitchRequest.Method = http.MethodGet
		channel := core.Channel{Request: twitchRequest}
		channelSummary := channel.RequestSummary()
		c.Context.DB.StoreChannelSummary(channelSummary)
	}
}

func (c Channel) updateSubscriptionSummary() {
	c.Loggger.Log("Aggregate on Subscription Summary...")
	for _, credential := range c.Context.Credentials {
		twitchRequest := core.NewUserAccessTokenRequest(c.OAuth2, credential.TokenResponse)
		twitchRequest.Logger = c.Loggger
		twitchRequest.Method = http.MethodGet
		cc := core.NewChannel(twitchRequest)
		lastChannelEntry := c.Context.DB.GetLastUpdatedChannelSummary(credential.ChannelName)
		s := cc.GetSubscriptionSummary(lastChannelEntry.IDTwitch)

		log.Println(fmt.Sprintf("Not used: %#v", s))
	}
}

// GetLastStreamSession retrieves last stream video information
func (c Channel) GetLastStreamSession() {
	channels := c.DB.GetChannels()
	for _, channel := range channels {
		channelAPI := core.Channel{
			Request: c.Request,
		}
		videos := channelAPI.GetVideosFromID(channel.IDTwitch, 100)
		for _, video := range videos {
			c.Context.DB.RegisterVideoToChannel(channel.IDTwitch, video)
		}
	}
}

// Process aggregation
func (c Channel) Process() {
	c.Loggger.Log("Start Channel aggregation...")
	c.updateChannelSummary()
	c.updateSubscriptionSummary()
	c.GetLastStreamSession()
}
