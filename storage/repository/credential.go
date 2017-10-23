package repository

import (
	"strconv"

	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// CredentialRepository handles credential database query
type CredentialRepository struct {
	*Repository
}

// NewCredentialRepository return a credential repository
func NewCredentialRepository(db *storage.Database, l logger.Logger) CredentialRepository {
	commonRepository := NewRepository(db, l)
	r := CredentialRepository{
		Repository: commonRepository,
	}

	return r
}

// getUID return a uid from the credential information
func (r *CredentialRepository) getUID(c model.Credential) string {
	return c.AppName + strconv.FormatInt(c.ChannelID, 10)
}

// SaveUserCredential saves or updates the credential
func (r *CredentialRepository) SaveUserCredential(c model.Credential) bool {
	uid := r.getUID(c)
	query := storage.Query{
		Query: `
			INSERT INTO ` + model.CredentialTable + `
			(uid, channel_name, channel_id, access_token, refresh_token, scope, expires_in, email)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE
				channel_name = ?,
				access_token = ?,
				refresh_token = ?,
				scope = ?,
				expires_in = ?,
				email = ?
		`,
		Parameters: map[string]interface{}{
			"uid":           uid,
			"channel_name":  c.ChannelName,
			"channel_id":    c.ChannelID,
			"access_token":  c.AccessToken,
			"refresh_token": c.RefreshToken,
			"scope":         c.Scopes,
			"expires_in":    c.ExpiresIn,
			"email":         c.Email,
		},
	}

	state := r.Database.Run(query,
		uid,
		c.ChannelName,
		c.ChannelID,
		c.AccessToken,
		c.RefreshToken,
		c.Scopes,
		c.ExpiresIn,
		c.Email,
		c.ChannelName,
		c.AccessToken,
		c.RefreshToken,
		c.Scopes,
		c.ExpiresIn,
		c.Email,
	)

	return state
}

// SaveAppCredential saves or updates the credential
func (r *CredentialRepository) SaveAppCredential(c model.Credential) bool {
	uid := r.getUID(c)
	query := storage.Query{
		Query: `
			INSERT INTO ` + model.CredentialTable + `
			(uid, app_name, access_token, refresh_token, scope, expires_in)
			VALUES(?, ?, ?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE
				access_token = ?,
				expires_in = ?
		`,
		Parameters: map[string]interface{}{
			"uid":           uid,
			"app_name":      c.AppName,
			"access_token":  c.AccessToken,
			"refresh_token": c.RefreshToken,
			"scope":         c.Scopes,
			"expires_in":    c.ExpiresIn,
		},
	}

	state := r.Database.Run(query,
		uid,
		c.AppName,
		c.AccessToken,
		c.RefreshToken,
		c.Scopes,
		c.ExpiresIn,
		c.AccessToken,
		c.RefreshToken,
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
		&credential.AccessToken,
		&credential.RefreshToken,
		&credential.Scopes,
		&credential.ExpiresIn,
		&credential.DateUpdated,
		&credential.Email,
	)

	return credential
}

// GetAppToken retrieve the credential token for a specific App
func (r CredentialRepository) GetAppToken(appName string) (model.Credential, bool) {
	query := storage.Query{
		Query: `
			SELECT
				id,
				app_name,
				access_token,
				scope,
				expires_in,
				date_updated
			FROM ` + model.CredentialTable + `
			WHERE app_name = ?
		`,
		Parameters: map[string]interface{}{
			"app_name": appName,
		},
	}

	var credential = model.Credential{}
	row := r.Database.QueryRow(query, appName)

	succeed := r.Database.ScanRow(row,
		&credential.ID,
		&credential.AppName,
		&credential.AccessToken,
		&credential.Scopes,
		&credential.ExpiresIn,
		&credential.DateUpdated,
	)

	return credential, succeed
}

// GetUserCredentials return a credentials list
func (r *CredentialRepository) GetUserCredentials() []model.Credential {
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
			WHERE
				app_name IS NULL
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
			&credential.Scopes,
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
