package aggregation

import (
	"log"

	"github.com/wonderstream/twitch/core"
)

// Channel aggregation contains requirement to handle the process
type Channel struct {
	Context
}

func (c Channel) updateChannelSummary() {
	log.Println("Aggregate on Channel Summary...")
	for _, credential := range c.Context.Credentials {
		twitchRequest := core.NewRequest(c.OAuth2, credential.TokenResponse)
		channel := core.Channel{Request: twitchRequest}
		channelSummary := channel.RequestSummary()
		c.Context.DB.StoreChannelSummary(channelSummary)
	}
}

func (c Channel) updateSubscriptionSummary() {
	log.Println("Aggregate on Subscription Summary...")
	for _, credential := range c.Context.Credentials {
		twitchRequest := core.NewRequest(c.OAuth2, credential.TokenResponse)
		cc := core.NewChannel(twitchRequest)
		lastChannelEntry := c.Context.DB.GetLastUpdatedChannelSummary(credential.ChannelName)
		s := cc.GetSubscriptionSummary(lastChannelEntry.IDTwitch)

		log.Println("NOT USED: ", s)
	}
}

// Process aggregation
func (c Channel) Process() {
	log.Println("Start Channel aggregation...")
	c.updateChannelSummary()
	c.updateSubscriptionSummary()
}
