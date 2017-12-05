package repository

import (
	"time"

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

// StoreCredential add new credential
func (r *CredentialRepository) StoreCredential(c model.Credential) bool {
	newCredential := model.Credential{}
	newCredential.MetaDateAdd = time.Now()
	c.MetaDateUpdate = time.Now()
	err := r.Database.Gorm.
		Where(model.Credential{ChannelID: c.ChannelID}).
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

// GetUserCredentials return a credentials list
func (r *CredentialRepository) GetUserCredentials() []model.Credential {
	credentials := []model.Credential{}
	r.Database.Gorm.Find(&credentials)

	return credentials
}
