package model

import (
	"time"
)

const (
	// ChannelTable database table
	ChannelTable = "api_channels"
)

// Channel mapping table
type Channel struct {
	MetaID               int64 `gorm:"primary_key:true"`
	MetaDateAdd          time.Time
	ID                   int64 `gorm:"index"`
	Name                 string
	Description          string `gorm:"type:TEXT"`
	Email                string
	URL                  string
	Videos               []Video `gorm:"ForeignKey:MetalChannelID;AssociationForeignKey:ID"`
	Mature               bool
	Status               string
	BroadcasterLanguage  string
	DisplayName          string
	Game                 string
	Language             string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Partner              bool
	Logo                 string
	VideoBanner          string
	ProfileBanner        string
	ProfileBannerBGColor string
	Views                int64
	Followers            int64
	BroadcasterType      string
	StreamKey            string
}

// TableName set be singular
func (Channel) TableName() string {
	return ChannelTable
}
