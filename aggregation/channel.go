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
	for _, credential := range c.Context.Credentials {
		log.Println("Getting credential for " + credential.ChannelName)
		twitchRequest := core.NewRequest(c.OAuth2, credential.TokenResponse)
		channel := core.Channel{Request: twitchRequest}
		channelSummary := channel.RequestSummary()

		c.Context.DB.StoreChannelSummary(channelSummary)
	}
}

// Process aggregation
func (c Channel) Process() {
	log.Println("Start Channel aggregation...")
	c.updateChannelSummary()
}
