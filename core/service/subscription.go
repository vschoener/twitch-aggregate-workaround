package service

import (
	"fmt"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
)

// SubscriptionService handles processes for the subscription
type SubscriptionService struct {
	*core.Request
}

// GetSubscription return the subscription summary of the channel ID
func (s SubscriptionService) GetSubscription(channelID int64) model.Subscription {
	subscription := model.Subscription{}
	url := fmt.Sprintf("%s/%d/subscriptions", core.ChannelsURI, channelID)
	s.SendRequest(url, &subscription)

	return subscription
}
