package service

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

// NewAuthService constructor
func NewAuthService(oauth2 *core.OAuth2, db *storage.Database) *Auth {
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

	twitchRequest := core.NewRequest(a.oauth2, &token)
	twitchRequest.Logger = logger.Share()
	twitchRequest.Logger.SetPrefix("LIBRARY")

	channelService := service.ChannelService{}
	userService := service.UserService{}
	credentialRepository := repository.NewCredentialRepository(a.db, logger)
	channelRepository := repository.NewChannelRepository(a.db, logger)
	userRepository := repository.NewUserRepository(a.db, logger)

	channel, err := channelService.GetInfo(twitchRequest)
	if err != nil {
		return err
	}

	user, err := userService.GetByName(channel.Name, twitchRequest)
	if err != nil {
		return err
	}

	sChannel := transformer.TransformCoreChannelToStorageChannel(channel)
	credential := transformer.TransformCoreTokenResponseToStorageCredential(token)
	sUser := transformer.TransformCoreUserToStorageUser(user)

	// TODO: Review channel ID Auth
	credential.ChannelID = sChannel.ID
	credential.ChannelName = sChannel.Name
	credential.Email = sChannel.Email
	if false == credentialRepository.StoreCredential(credential) {
		return errors.New("Error getting credential")
	}
	if false == channelRepository.StoreChannel(sChannel) {
		return errors.New("Error storing channel")
	}
	if false == userRepository.StoreUser(sUser) {
		return errors.New("Error storing user")
	}

	return nil
}
