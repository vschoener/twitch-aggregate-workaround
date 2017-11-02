package repository

import (
	"github.com/wonderstream/twitch/storage/model"
)

// ChannelVideoRepository handles channel video database query
type ChannelVideoRepository struct {
	*Repository
}

// RegisterVideoToChannel inserts or updates video information to the channel ID
func (r ChannelVideoRepository) RegisterVideoToChannel(channelID int64, video model.ChannelVideo) bool {
	video.ChannelID = channelID
	newVideo := model.ChannelVideo{}
	err := r.Database.Gorm.
		Where(model.ChannelVideo{VideoID: video.VideoID}).
		Assign(video).
		FirstOrCreate(&newVideo).
		Error

	return r.CheckErr(err)
}
