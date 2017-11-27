package repository

import (
	"time"

	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// ChannelRepository handles channel database query
type ChannelRepository struct {
	*Repository
}

// NewChannelRepository return a credential repository
func NewChannelRepository(db *storage.Database, l logger.Logger) ChannelRepository {
	commonRepository := NewRepository(db, l)
	r := ChannelRepository{
		Repository: commonRepository,
	}

	return r
}

// StoreChannel will add new entry everytime to have an history
func (r ChannelRepository) StoreChannel(channel model.Channel) bool {
	channel.MetaDateAdd = time.Now()
	err := r.Database.Gorm.Create(&channel).Error

	return r.CheckErr(err)
}

// GetLastChannelsRecord return all channel stored in the database
func (r ChannelRepository) GetLastChannelsRecord() []model.Channel {
	channels := []model.Channel{}
	r.Database.Gorm.
		Group("channel_id").
		Order("id DESC").
		Find(&channels)

	return channels
}

// GetLastRecorded returns the last recorded summary from Database
func (r ChannelRepository) GetLastRecorded(channelName string) (model.Channel, bool) {
	channel := model.Channel{}
	found := !r.Database.Gorm.
		Where("name = ?", channelName).
		Order("id DESC").
		Find(&channel).
		RecordNotFound()

	return channel, found
}
