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
	ID             int64
	UID            string `gorm:"column:uid"`
	MetaDateAdd    time.Time
	MetaDateUpdate time.Time
	AppName        string
	ChannelName    string
	ChannelID      int64
	Email          string
	AccessToken    string
	RefreshToken   string
	ExpiresIn      int64
	Scopes         string
}

// IsSet is a shortcut function to know if the credential is Found or Set
func (c Credential) IsSet() bool {
	return c.ID > 0
}
