package repository

import (
	"log"

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
func (r ChannelVideoRepository) GetAirTime(channelID int64, queryFilter storage.QueryFilter) int64 {
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

	queryFilter.DateField = "recorded_at"
	r.applyFilter(db, queryFilter)

	err := db.Scan(&result).Error

	if nil != err {
		log.Println(err)
	}

	return result.Total
}

// StreamedGame gives information about how much time a game has been played
type StreamedGame struct {
	Name        string
	TotalPlayed int64
}

// GetGames played
func (r ChannelVideoRepository) GetGames(channelID int64, queryFilter storage.QueryFilter) []StreamedGame {
	games := []StreamedGame{}

	db := r.Database.Gorm.
		Model(&model.ChannelVideo{}).
		Select("game as Name, Count(*) TotalPlayed").
		Where(`channel_id = ?
			AND broadcast_type = ?`,
			channelID,
			"archive",
		).Order("TotalPlayed DESC").
		Group("game")

	queryFilter.DateField = "recorded_at"
	r.applyFilter(db, queryFilter)

	rows, err := db.Rows()
	defer rows.Close()

	if nil != err {
		log.Println(err)
	}

	// Would prefer to scan the list instead of browsing and store the row but
	// it was not working :(
	for rows.Next() {
		t := StreamedGame{}
		rows.Scan(&t.Name, &t.TotalPlayed)
		games = append(games, t)
	}

	return games
}
