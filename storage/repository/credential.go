package repository

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// CredentialRepository handles credential database query
type CredentialRepository struct {
	*Repository
}

// SaveCredential saves or updates the credential
func (r *CredentialRepository) SaveCredential(cs model.Channel, token core.TokenResponse) bool {
	query := storage.Query{
		Query: `
			INSERT INTO ` + model.CredentialTable + `
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

	state := r.Database.Run(query,
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

	return state
}

// GetCredential will retrieve the oauth2 token information returning a TokenResponse
// as reference
func (r CredentialRepository) GetCredential(channelName string) model.Credential {
	query := storage.Query{
		Query: `
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
			FROM ` + model.CredentialTable + `
			WHERE channel_name = ?
		`,
		Parameters: map[string]interface{}{
			"channel_name": channelName,
		},
	}

	var credential = model.Credential{}
	row := r.Database.QueryRow(query, channelName)
	r.Database.ScanRow(row,
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

	return credential
}

// GetCredentials return a credentials list
func (r *CredentialRepository) GetCredentials() []model.Credential {
	query := storage.Query{
		Query: `
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
			FROM ` + model.CredentialTable + `
		`,
	}

	rows := r.Database.Query(query)
	if rows == nil {
		return nil
	}

	credentials := []model.Credential{}
	defer rows.Close()

	for rows.Next() {
		credential := model.Credential{}
		state := r.Database.ScanRows(rows,
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
		if state {
			credentials = append(credentials, credential)
		}
	}
	if err := rows.Err(); err != nil {
		r.Logger.LogInterface(err)
	}

	return credentials
}
