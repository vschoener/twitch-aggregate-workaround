package service

import (
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/repository"
)

// ChannelService handle needs for app / aggreation
type ChannelService struct {
}

// NewChannelService constructor
func NewChannelService() *ChannelService {
	return &ChannelService{}
}

// GetPrimaryGame returns the game mostly played
func (s ChannelService) GetPrimaryGame(channelID int64, f storage.QueryFilter, cvr repository.ChannelVideoRepository) repository.StreamedGame {
	primaryGameStreamed := repository.StreamedGame{}
	games := cvr.GetGames(channelID, f)

	for _, game := range games {
		if primaryGameStreamed.TotalPlayed < game.TotalPlayed {
			primaryGameStreamed = game
		}
	}
	return primaryGameStreamed
}
