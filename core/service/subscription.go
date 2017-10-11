package service

import (
	"fmt"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
)

// SubscriptionService handles processes for the subscription
type SubscriptionService struct {
}

// NewSubscriptionService constructor
func NewSubscriptionService() SubscriptionService {
	return SubscriptionService{}
}

// GetSubscription return the subscription summary of the channel ID
func (s SubscriptionService) GetSubscription(channelID int64, r *core.Request) model.Subscription {
	subscription := model.Subscription{}
	url := fmt.Sprintf("%s/%d/subscriptions", core.ChannelsURI, channelID)
	r.SendRequest(url, &subscription)

	return subscription
}
