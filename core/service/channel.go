package service

import (
	"fmt"

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
func (s ChannelService) GetInfo(r *core.Request) (model.Channel, error) {
	channel := model.Channel{}
	err := r.SendRequest(core.ChannelURI, &channel)

	return channel, err
}

// GetInfoByID return public channel information
func (s ChannelService) GetInfoByID(id int64, r *core.Request) (model.Channel, error) {
	channel := model.Channel{}
	err := r.SendRequest(fmt.Sprintf("%s/%d", core.ChannelsURI, id), &channel)

	return channel, err
}
