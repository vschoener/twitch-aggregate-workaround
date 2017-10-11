package aggregation

import (
	"net/http"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// Auth aggregation handler
type Auth struct {
	oauth2 *core.OAuth2
	db     *storage.Database
}

// NewAuthAggregation constructor
func NewAuthAggregation(oauth2 *core.OAuth2, db *storage.Database) *Auth {
	return &Auth{
		oauth2: oauth2,
		db:     db,
	}
}

// HandleHTTPRequest process the request coming from Twitch to save the
// the information into the storage.
// This function should be reach one the use has accepted the app rights
func (a *Auth) HandleHTTPRequest(w http.ResponseWriter, twRequest *http.Request, logger logger.Logger) error {
	token, err := a.oauth2.RequestToken(twRequest)

	if err != nil {
		return err
	}

	twitchRequest := core.NewUserAccessTokenRequest(a.oauth2, token)
	twitchRequest.Logger = logger

	channelService := service.ChannelService{}

	channel := channelService.GetInfo(twitchRequest)
	commonRepository := repository.NewRepository(a.db, logger)
	credentialRepository := repository.CredentialRepository{
		Repository: commonRepository,
	}
	channelRepository := repository.ChannelRepository{
		Repository: commonRepository,
	}

	sChannel := transformer.TransformCoreChannelToStorageChannel(channel)
	credentialRepository.SaveCredential(sChannel, token)
	channelRepository.StoreChannel(sChannel)

	return nil
}
