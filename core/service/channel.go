package service

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
)

// ChannelService handles processes for the channel
type ChannelService struct {
}

// NewChannelService constructor
func NewChannelService() ChannelService {
	return ChannelService{}
}

// GetInfo retrieve information
func (s ChannelService) GetInfo(r *core.Request) model.Channel {
	channel := model.Channel{}
	r.SendRequest(core.ChannelURI, &channel)

	return channel
}
