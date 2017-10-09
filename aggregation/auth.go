package aggregation

import (
	"net/http"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
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

	channel := core.Channel{Request: twitchRequest}
	channelSummary := channel.RequestSummary()
	a.db.RecordToken(channelSummary, token)

	return nil
}
