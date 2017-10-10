package service

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
)

// ChannelService handles processes for the channel
type ChannelService struct {
	*core.Request
}

// GetInfo retrieve information
func (s ChannelService) GetInfo() model.Channel {
	channel := model.Channel{}
	s.SendRequest(core.ChannelURI, &channel)

	return channel
}
