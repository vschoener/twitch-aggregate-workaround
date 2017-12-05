package repository

import (
	"log"
	"time"

	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// VideoRepository handles channel video database query
type VideoRepository struct {
	*Repository
}

// NewChannelVideoRepository return a credential repository
func NewChannelVideoRepository(db *storage.Database, l logger.Logger) VideoRepository {
	commonRepository := NewRepository(db, l)
	r := VideoRepository{
		Repository: commonRepository,
	}

	return r
}

// RegisterVideoToChannel inserts or updates video information to the channel ID
func (r VideoRepository) RegisterVideoToChannel(channelID int64, videos []model.Video) bool {

	for _, video := range videos {
		video.MetaChannelID = channelID
		newVideo := model.Video{}
		newVideo.MetaDateAdd = time.Now()
		err := r.Database.Gorm.
			Where(model.Video{ID: video.ID}).
			Assign(video).
			FirstOrCreate(&newVideo).
			Error

		if err != nil {
			r.CheckErr(err)
		}
	}

	return true
}

// GetAirTime returns the total stream in seconds
func (r VideoRepository) GetAirTime(channelID int64, queryFilter storage.QueryFilter) int64 {
	type Result struct {
		Total int64
	}

	result := Result{}
	db := r.Database.Gorm.
		Model(&model.Video{}).
		Select("SUM(length) total").
		Where(`meta_channel_id = ?
			AND broadcast_type = ?`,
			channelID,
			"archive",
		).Group("meta_channel_id")

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
func (r VideoRepository) GetGames(channelID int64, queryFilter storage.QueryFilter) []StreamedGame {
	games := []StreamedGame{}

	db := r.Database.Gorm.
		Model(&model.Video{}).
		Select("game as Name, Count(*) TotalPlayed").
		Where(`meta_channel_id = ?
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
