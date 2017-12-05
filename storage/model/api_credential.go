package model

import (
	"time"
)

const (
	// CredentialTable database table
	CredentialTable = "api_credentials"
)

// Credential map the database table
type Credential struct {
	ID             int64
	Channel        Channel
	MetaDateAdd    time.Time
	MetaDateUpdate time.Time
	ChannelName    string
	ChannelID      int64 `gorm:"unique_index"`
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

// TableName set be singular
func (Credential) TableName() string {
	return CredentialTable
}
