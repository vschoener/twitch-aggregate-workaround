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

// SaveCredential saves or updates the credential
func (s *Database) SaveCredential(cs core.ChannelSummary, token core.TokenResponse) error {
	queryLogger := QueryLogger{
		Query: `
			INSERT INTO ` + credentialTable + `
			(channel_name, channel_id, access_token, refresh_token, scope, expires_in, email)
			VALUES(?, ?, ?, ?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE
				channel_name = ?,
				access_token = ?,
				refresh_token = ?,
				scope = ?,
				expires_in = ?,
				email = ?
		`,
		Parameters: map[string]interface{}{
			"channel_name":  cs.Name,
			"channel_id":    cs.IDTwitch,
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
			"scope":         token.Scope,
			"expires_in":    token.ExpiresIn,
			"email":         cs.Email,
		},
	}

	s.Logger.LogInterface(queryLogger)
	stmt, err := s.DB.Prepare(queryLogger.Query)

	if err != nil {
		s.Logger.Log(err.Error())
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		cs.Name,
		cs.IDTwitch,
		token.AccessToken,
		token.RefreshToken,
		token.Scope,
		token.ExpiresIn,
		cs.Email,
		cs.Name,
		token.AccessToken,
		token.RefreshToken,
		token.Scope,
		token.ExpiresIn,
		cs.Email,
	)

	if err != nil {
		s.Logger.Log(err.Error())
		return err
	}

	return nil
}

// GetCredential will retrieve the oauth2 token information returning a TokenResponse
// as reference
func (s *Database) GetCredential(channelName string) Credential {

	var credential = Credential{}

	err := s.DB.QueryRow(`
		SELECT
			id,
			channel_name,
			channel_id,
			access_token,
			refresh_token,
			scope,
			expires_in,
			date_updated,
			email
		FROM `+credentialTable+`
		WHERE channel_name = ?
	`, channelName).Scan(
		&credential.ID,
		&credential.ChannelName,
		&credential.ChannelID,
		&credential.TokenResponse.AccessToken,
		&credential.TokenResponse.RefreshToken,
		&credential.TokenResponse.Scope,
		&credential.TokenResponse.ExpiresIn,
		&credential.DateUpdated,
		&credential.Email,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return credential
}

// GetCredentials return a credentials list
func (s *Database) GetCredentials() []Credential {
	rows, err := s.DB.Query(`
        SELECT
            id,
			channel_name,
			channel_id,
			access_token,
			refresh_token,
			scope,
			expires_in,
			date_updated,
			email
        FROM ` + credentialTable + `
    `)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	credentials := []Credential{}
	for rows.Next() {
		credential := Credential{}
		err := rows.Scan(
			&credential.ID,
			&credential.ChannelName,
			&credential.ChannelID,
			&credential.AccessToken,
			&credential.RefreshToken,
			&credential.Scope,
			&credential.ExpiresIn,
			&credential.DateUpdated,
			&credential.Email,
		)

		if err != nil {
			log.Fatal(err)
		}
		credentials = append(credentials, credential)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return credentials
}
