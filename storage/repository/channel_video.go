package repository

import (
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// ChannelVideoRepository handles channel video database query
type ChannelVideoRepository struct {
	*Repository
}

// RegisterVideoToChannel inserts or updates video information to the channel ID
func (r ChannelVideoRepository) RegisterVideoToChannel(channelID int64, video model.ChannelVideo) bool {
	query := storage.Query{
		Query: `
            INSERT INTO ` + model.ChannelVideoTable + `
            (
				channel_id,
                video_id,
                title,
                description,
                description_html,
                broadcast_id,
                broadcast_type,
                status,
                tag_list,
                views,
                url,
				language,
                created_at,
                viewable,
                viewable_at,
                published_at,
                recorded_at,
                game,
                length
            )
            VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
            ON DUPLICATE KEY UPDATE
                id = id
        `,
		Parameters: map[string]interface{}{
			"channelID":    channelID,
			"channelVideo": video,
		},
	}

	state := r.Database.Run(query,
		channelID,
		video.VideoID,
		video.Title,
		video.Description,
		video.DescriptionHTML,
		video.BrodcastID,
		video.BrodcastType,
		video.Status,
		video.TagList,
		video.Views,
		video.URL,
		video.Language,
		video.CreatedAt,
		video.Viewable,
		video.ViewableAt,
		video.PublishedAt,
		video.RecordedAt,
		video.Game,
		video.Length,
	)

	return state
}
