package storage

import (
	"fmt"
	"time"

	"github.com/wonderstream/twitch/core"
)

const (
	channelTable      = "channel"
	channelVideoTable = "channel_video"
)

// Channel mapping table
type Channel struct {
	ID      int64
	DateAdd time.Time
	core.ChannelSummary
}

// ChannelVideo mapping table
type ChannelVideo struct {
	ID        int64
	ChannelID int64
	DateAdd   time.Time
	core.ChannelVideo
}

// RegisterVideoToChannel inserts or updates video information to the channel ID
func (d *Database) RegisterVideoToChannel(channelID int64, video core.ChannelVideo) bool {
	queryLogger := QueryLogger{
		Query: `
            INSERT INTO ` + channelVideoTable + `
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

	d.Logger.Log(fmt.Sprintf("RegisterVideoToChannel on %#v", queryLogger))
	stmt, err := d.DB.Prepare(queryLogger.Query)

	if err != nil {
		d.Logger.LogInterface(err)
		return false
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		channelID,
		video.ID,
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

	if err != nil {
		d.Logger.LogInterface(err)
		return false
	}

	return true
}

// StoreChannelSummary will add new entry everytime to have an history
func (d *Database) StoreChannelSummary(channelSummary core.ChannelSummary) bool {
	queryLogger := QueryLogger{
		Query: `
	        INSERT INTO ` + channelTable + `
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
			"ChannelSummary": channelSummary,
		},
	}

	stmt, err := d.DB.Prepare(queryLogger.Query)

	if err != nil {
		d.Logger.LogInterface(err)
		return false
	}

	defer stmt.Close()
	d.Logger.Log(fmt.Sprintf("StoreChannelSummary on %#v", queryLogger))
	_, err = stmt.Exec(
		channelSummary.Mature,
		channelSummary.Status,
		channelSummary.BroadcasterLanguage,
		channelSummary.DisplayName,
		channelSummary.Game,
		channelSummary.Language,
		channelSummary.IDTwitch,
		channelSummary.Name,
		channelSummary.CreatedAt,
		channelSummary.UpdatedAt,
		channelSummary.Partner,
		channelSummary.Logo,
		channelSummary.VideoBanner,
		channelSummary.ProfileBanner,
		channelSummary.ProfileBannerBGColor,
		channelSummary.URL,
		channelSummary.Views,
		channelSummary.Followers,
		channelSummary.BroadcasterType,
		channelSummary.StreamKey,
		channelSummary.Email,
	)

	if err != nil {
		d.Logger.LogInterface(err)
		return false
	}

	return true
}

// GetChannels return all channel stored in the database
func (d *Database) GetChannels() []Channel {
	queryLogger := QueryLogger{
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
	        FROM ` + channelTable + `
			GROUP BY _id
			ORDER BY id DESC
	    `,
	}

	d.Logger.Log(fmt.Sprintf("StoreChannelSummary on %#v", queryLogger))
	rows, err := d.DB.Query(queryLogger.Query)

	if err != nil {
		d.Logger.LogInterface(err)
		return nil
	}

	defer rows.Close()
	channels := []Channel{}

	for rows.Next() {
		channel := Channel{}
		err := rows.Scan(
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
		if err != nil {
			d.Logger.LogInterface(err)
			return nil
		}

		d.Logger.LogInterface(rows)
		channels = append(channels, channel)
	}
	if err := rows.Err(); err != nil {
		d.Logger.LogInterface(err)
		return nil
	}

	return channels
}

// GetLastUpdatedChannelSummary returns the last recorded summary from Database
func (d *Database) GetLastUpdatedChannelSummary(channelName string) Channel {
	channel := Channel{}
	query := `
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
        FROM ` + channelTable + `
		WHERE name=?
		ORDER BY id DESC
		LIMIT 1
	`
	row := d.DB.QueryRow(query, channelName)

	d.Logger.Log(fmt.Sprintf("GetLastUpdatedChannelSummary on %#v", QueryLogger{
		Query: query,
		Parameters: map[string]interface{}{
			"name": channelName,
		},
	}))

	err := row.Scan(
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

	if err != nil {
		d.Logger.LogInterface(err)
	}

	return channel
}
