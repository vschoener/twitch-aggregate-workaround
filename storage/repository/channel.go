package repository

import (
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// ChannelRepository handles channel database query
type ChannelRepository struct {
	*Repository
}

// StoreChannel will add new entry everytime to have an history
func (r ChannelRepository) StoreChannel(channel model.Channel) bool {
	query := storage.Query{
		Query: `
	        INSERT INTO ` + model.ChannelTable + `
	        (
	            mature,
	            status,
	            broadcaster_language,
	            display_name,
	            game,
	            language,
	            _id,
	            name,
	            created_at,
	            updated_at,
	            partner,
	            logo,
	            video_banner,
	            profile_banner,
	            profile_banner_background_color,
	            url,
	            views,
	            followers,
	            broadcaster_type,
	            stream_key,
	            email
	        )
	        VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		Parameters: map[string]interface{}{
			"ChannelSummary": channel,
		},
	}

	state := r.Database.Run(query,
		channel.Mature,
		channel.Status,
		channel.BroadcasterLanguage,
		channel.DisplayName,
		channel.Game,
		channel.Language,
		channel.IDTwitch,
		channel.Name,
		channel.CreatedAt,
		channel.UpdatedAt,
		channel.Partner,
		channel.Logo,
		channel.VideoBanner,
		channel.ProfileBanner,
		channel.ProfileBannerBGColor,
		channel.URL,
		channel.Views,
		channel.Followers,
		channel.BroadcasterType,
		channel.StreamKey,
		channel.Email)

	return state
}

// GetChannels return all channel stored in the database
func (r ChannelRepository) GetChannels() []model.Channel {
	query := storage.Query{
		Query: `
	        SELECT
	            id,
	            mature,
	            status,
	            broadcaster_language,
	            display_name,
	            game,
	            language,
	            _id,
	            name,
	            created_at,
	            updated_at,
	            partner,
	            logo,
	            video_banner,
	            profile_banner,
	            profile_banner_background_color,
	            url,
	            views,
	            followers,
	            broadcaster_type,
	            stream_key,
	            email,
	            date_add
	        FROM ` + model.ChannelTable + `
			GROUP BY _id
			ORDER BY id DESC
	    `,
	}

	rows := r.Database.Query(query)
	if rows == nil {
		return nil
	}

	defer rows.Close()
	channels := []model.Channel{}

	for rows.Next() {
		channel := model.Channel{}
		state := r.Database.ScanRows(rows,
			&channel.ID,
			&channel.Mature,
			&channel.Status,
			&channel.BroadcasterLanguage,
			&channel.DisplayName,
			&channel.Game,
			&channel.Language,
			&channel.IDTwitch,
			&channel.Name,
			&channel.CreatedAt,
			&channel.UpdatedAt,
			&channel.Partner,
			&channel.Logo,
			&channel.VideoBanner,
			&channel.ProfileBanner,
			&channel.ProfileBannerBGColor,
			&channel.URL,
			&channel.Views,
			&channel.Followers,
			&channel.BroadcasterType,
			&channel.StreamKey,
			&channel.Email,
			&channel.DateAdd,
		)
		if state {
			channels = append(channels, channel)
		}
	}
	if err := rows.Err(); err != nil {
		r.Logger.LogInterface(err)
		return nil
	}

	return channels
}

// GetLastUpdatedChannelSummary returns the last recorded summary from Database
func (r ChannelRepository) GetLastUpdatedChannelSummary(channelName string) model.Channel {
	query := storage.Query{
		Query: `
	        SELECT
	            id,
	            mature,
	            status,
	            broadcaster_language,
	            display_name,
	            game,
	            language,
	            _id,
	            name,
	            created_at,
	            updated_at,
	            partner,
	            logo,
	            video_banner,
	            profile_banner,
	            profile_banner_background_color,
	            url,
	            views,
	            followers,
	            broadcaster_type,
	            stream_key,
	            email,
	            date_add
	        FROM ` + model.ChannelTable + `
			WHERE name=?
			ORDER BY id DESC
			LIMIT 1
		`,
		Parameters: map[string]interface{}{
			"name": channelName,
		},
	}
	row := r.Database.QueryRow(query, channelName)
	channel := model.Channel{}

	r.Database.ScanRow(row,
		&channel.ID,
		&channel.Mature,
		&channel.Status,
		&channel.BroadcasterLanguage,
		&channel.DisplayName,
		&channel.Game,
		&channel.Language,
		&channel.IDTwitch,
		&channel.Name,
		&channel.CreatedAt,
		&channel.UpdatedAt,
		&channel.Partner,
		&channel.Logo,
		&channel.VideoBanner,
		&channel.ProfileBanner,
		&channel.ProfileBannerBGColor,
		&channel.URL,
		&channel.Views,
		&channel.Followers,
		&channel.BroadcasterType,
		&channel.StreamKey,
		&channel.Email,
		&channel.DateAdd,
	)

	return channel
}
