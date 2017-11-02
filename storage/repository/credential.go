package repository

import (
	"strconv"
	"strings"

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
	if len(c.AppName) > 0 {
		return strings.ToUpper(c.AppName)
	}

	return strconv.FormatInt(c.ChannelID, 10)
}

// StoreCredential add new credential
func (r *CredentialRepository) StoreCredential(c model.Credential) bool {
	uid := r.getUID(c)
	newCredential := model.Credential{}
	err := r.Database.Gorm.
		Where(model.Credential{UID: uid}).
		Assign(c).
		FirstOrCreate(&newCredential).
		Error

	return r.CheckErr(err)
}

// GetCredential will retrieve the oauth2 token information returning a TokenResponse
// as reference
func (r CredentialRepository) GetCredential(channelName string) (model.Credential, bool) {
	var credential = model.Credential{}
	notFound := r.Database.Gorm.
		Where("channel_name = ?", channelName).
		Find(&credential).
		RecordNotFound()

	return credential, !notFound
}

// GetAppToken retrieve the credential token for a specific App
func (r CredentialRepository) GetAppToken(appName string) (model.Credential, bool) {
	var credential = model.Credential{}
	notFound := r.Database.Gorm.
		Where("app_name = ?", appName).
		Find(&credential).
		RecordNotFound()

	return credential, !notFound
}

// GetUserCredentials return a credentials list
func (r *CredentialRepository) GetUserCredentials() []model.Credential {
	credentials := []model.Credential{}
	r.Database.Gorm.Where("app_name IS NULL").Find(&credentials)

	return credentials
}
