package aggregation

import (
	"errors"
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

// HandleUserAccessTokenHTTPRequest process the request coming from Twitch to save the
// the information into the storage.
// This function should be reach one the use has accepted the app rights
func (a *Auth) HandleUserAccessTokenHTTPRequest(w http.ResponseWriter, twRequest *http.Request, logger logger.Logger) error {
	authorizationCode := twRequest.URL.Query().Get("code")
	if len(authorizationCode) == 0 {
		err := errors.New("The query 'code' parameter is missing, please try again or contact us at contact@wonderstream.tv")
		if twError := twRequest.URL.Query().Get("error"); len(twError) > 0 {
			err = errors.New(twRequest.URL.Query().Get("error_description"))
		}
		return err
	}
	token, err := a.oauth2.RequestUserAcccessToken(authorizationCode)

	if err != nil {
		return err
	}

	if len(token.AccessToken) <= 0 {
		return errors.New("Token empty, please ask for a new authorization code")
	}

	twitchRequest := core.NewUserAccessTokenRequest(a.oauth2, token)
	twitchRequest.Logger = logger.Share()
	twitchRequest.Logger.SetPrefix("LIBRARY")

	channelService := service.ChannelService{}
	commonRepository := repository.NewRepository(a.db, logger)
	credentialRepository := repository.CredentialRepository{
		Repository: commonRepository,
	}
	channelRepository := repository.ChannelRepository{
		Repository: commonRepository,
	}

	channel := channelService.GetInfo(twitchRequest)
	sChannel := transformer.TransformCoreChannelToStorageChannel(channel)
	credential := transformer.TransformCoreTokenResponseToStorageCredential(token)
	credential.ChannelID = sChannel.IDTwitch
	credential.ChannelName = sChannel.Name
	credential.Email = sChannel.Email
	credentialRepository.SaveUserCredential(credential)
	channelRepository.StoreChannel(sChannel)

	return nil
}
