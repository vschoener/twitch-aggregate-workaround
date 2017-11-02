package repository

import (
	"log"
	"time"

	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// ChannelVideoRepository handles channel video database query
type ChannelVideoRepository struct {
	*Repository
}

// NewChannelVideoRepository return a credential repository
func NewChannelVideoRepository(db *storage.Database, l logger.Logger) ChannelVideoRepository {
	commonRepository := NewRepository(db, l)
	r := ChannelVideoRepository{
		Repository: commonRepository,
	}

	return r
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

// GetAirTime returns the total stream in seconds
func (r ChannelVideoRepository) GetAirTime(channelID int64, from *time.Time, to *time.Time) int64 {
	type Result struct {
		Total int64
	}

	result := Result{}
	db := r.Database.Gorm.
		Model(&model.ChannelVideo{}).
		Select("SUM(length) total").
		Where(`channel_id = ?
			AND broadcast_type = ?`,
			channelID,
			"archive",
		).Group("channel_id")

	if nil != from {
		db = db.Where("recorded_at >= ?", from.Format("2006-01-02"))
	}
	if nil != to {
		db = db.Where("recorded_at <= DATE_ADD(?, INTERVAL 1 DAY)", to.Format("2006-01-02"))
	}

	err := db.Scan(&result).Error

	if nil != err {
		log.Println(err)
	}

	return result.Total
}
