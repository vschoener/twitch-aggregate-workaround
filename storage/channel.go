package storage

import (
	"log"
	"time"

	"github.com/wonderstream/twitch/core"
)

const (
	channelTable = "channel"
)

// Channel mapping table
type Channel struct {
	ID      int64
	DateAdd time.Time
	core.ChannelSummary
}

// StoreChannelSummary will add new entry everytime to have an history
func (s *Database) StoreChannelSummary(channelSummary core.ChannelSummary) bool {
	stmt, err := s.DB.Prepare(`
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
        VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

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
		log.Fatal(err)
	}

	return true
}

// GetChannels return all channel stored in the database
func (s *Database) GetChannels() []Channel {
	rows, err := s.DB.Query(`
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
        FROM `+channelTable+`
    `, nil)

	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
		channels = append(channels, channel)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return channels
}
