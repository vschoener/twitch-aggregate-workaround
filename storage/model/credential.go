package model

import (
	"time"
)

const (
	// CredentialTable database table
	CredentialTable = "credentials"
)

// Credential map the database table
type Credential struct {
	UID          string `gorm:"column:uid"`
	ID           int64
	AppName      string
	ChannelName  string
	ChannelID    int64
	Email        string
	DateUpdated  time.Time
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	Scopes       string
}

// IsSet is a shortcut function to know if the credential is Found or Set
func (c Credential) IsSet() bool {
	return c.ID > 0
}
