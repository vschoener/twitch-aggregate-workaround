package storage

import (
	"database/sql"
	"log"
	"time"

	"github.com/wonderstream/twitch/core"
)

const (
	credentialTable = "credential"
)

// Credential map the database table
type Credential struct {
	core.TokenResponse
	ID          int64
	ChannelName string
	DateUpdated time.Time
}

// IsSet is a shortcut function to know if the credential is Found or Set
func (c Credential) IsSet() bool {
	return len(c.ChannelName) > 0
}

// Add new credential, used by RecordToken
func (s *Database) insertCredential(cs core.ChannelSummary, token core.TokenResponse) {
	stmt, err := s.DB.Prepare(`
		INSERT INTO ` + credentialTable + `
		(channelName, access_token, refresh_token, scope, expires_in)
		VALUES(?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		cs.Name,
		token.AccessToken,
		token.RefreshToken,
		token.Scope,
		token.ExpiresIn,
	)

	if err != nil {
		log.Fatal(err)
	}
}

// Update credential, used by RecordToken
func (s *Database) updateCredential(cs core.ChannelSummary, token core.TokenResponse) {
	stmt, err := s.DB.Prepare(`
		UPDATE ` + credentialTable + ` SET
			access_token=?,
			refresh_token=?,
			scope=?,
			expires_in=?,
			date_updated=NOW()
		WHERE channelName=?
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		token.AccessToken,
		token.RefreshToken,
		token.Scope,
		token.ExpiresIn,
		cs.Name,
	)

	if err != nil {
		log.Fatal(err)
	}
}

// RecordToken used to save token information inside the database
// If any error occure, log.Fatal is executed
func (s *Database) RecordToken(cs core.ChannelSummary, token core.TokenResponse) {

	credential := s.GetToken(cs.Name)

	if credential.IsSet() {
		s.updateCredential(cs, token)
	} else {
		s.insertCredential(cs, token)
	}
}

// GetToken will retrieve the oauth2 token information returning a TokenResponse
// as reference
func (s *Database) GetToken(channelName string) Credential {

	var credential = Credential{}

	err := s.DB.QueryRow(`
		SELECT
			id,
			channelName,
			access_token,
			refresh_token,
			scope,
			expires_in,
			date_updated
		FROM `+credentialTable+`
		WHERE channelName = ?
	`, channelName).Scan(
		&credential.ID,
		&credential.ChannelName,
		&credential.TokenResponse.AccessToken,
		&credential.TokenResponse.RefreshToken,
		&credential.TokenResponse.Scope,
		&credential.TokenResponse.ExpiresIn,
		&credential.DateUpdated,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return credential
}
