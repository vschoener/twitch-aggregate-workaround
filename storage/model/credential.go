package model

import (
	"time"

	"github.com/wonderstream/twitch/core"
)

const (
	// CredentialTable database table
	CredentialTable = "credential"
)

// Credential map the database table
type Credential struct {
	ID          int64
	ChannelName string
	ChannelID   int64
	Email       string
	DateUpdated time.Time
	core.TokenResponse
}

// IsSet is a shortcut function to know if the credential is Found or Set
func (c Credential) IsSet() bool {
	return len(c.ChannelName) > 0
}
