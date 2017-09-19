package aggregation

import (
	"net/http"

	"github.com/wonderstream/twitch/core"
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
// the information into the storage
func (a *Auth) HandleHTTPRequest(w http.ResponseWriter, req *http.Request) {
	token, err := a.oauth2.RequestToken(req)

	if err != nil {
		http.Redirect(w, req, a.oauth2.ErrorRedirectURL, http.StatusUnauthorized)
		return
	}

	twitchRequest := core.NewRequest(a.oauth2, token)
	channel := core.Channel{Request: twitchRequest}
	channelSummary := channel.RequestSummary()
	a.db.RecordToken(channelSummary, token)
}
