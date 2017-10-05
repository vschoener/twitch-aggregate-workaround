package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// OAuth2 manager
type OAuth2 struct {
	Scopes map[string]string
	*TwitchSettings
}

// TokenResponse use to store response body when getting a new token
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int64  `json:"expires_in"`
}

const (
	// ScopeUserRead provides read access to non-public user information, such
	// as their email address.
	ScopeUserRead = "user_read"
	// ScopeUserBlocksEdit provides the ability to ignore or unignore on
	// behalf of a user.
	ScopeUserBlocksEdit = "user_blocks_edit"
	// ScopeUserBlocksRead provides read access to a user's list of ignored
	// users.
	ScopeUserBlocksRead = "user_blocks_read"
	// ScopeUserFollowsEdit provides access to manage a user's followed
	// channels.
	ScopeUserFollowsEdit = "user_follows_edit"
	// ScopeChannelRead provides read access to non-public channel information,
	// including email address and stream key.
	ScopeChannelRead = "channel_read"
	// ScopeChannelEditor provides write access to channel metadata (game,
	// status, etc).
	ScopeChannelEditor = "channel_editor"
	// ScopeChannelCommercial provides access to trigger commercials on
	// channel.
	ScopeChannelCommercial = "channel_commercial"
	// ScopeChannelStream provides the ability to reset a channel's stream key.
	ScopeChannelStream = "channel_stream"
	// ScopeChannelSubscriptions provides read access to all subscribers to
	// your channel.
	ScopeChannelSubscriptions = "channel_subscriptions"
	// ScopeUserSubscriptions provides read access to subscriptions of a user.
	ScopeUserSubscriptions = "user_subscriptions"
	// ScopeChannelCheckSubscription provides read access to check if a user is
	// subscribed to your channel.
	ScopeChannelCheckSubscription = "channel_check_subscription"
	// ScopeChatLogin provides the ability to log into chat and send messages.
	ScopeChatLogin = "chat_login"
)

// NewOAuth2 constructor
func NewOAuth2(ts *TwitchSettings) *OAuth2 {
	return &OAuth2{
		TwitchSettings: ts,
	}
}

// RequestToken has to request a token using Query "code" paramater available
// inside the twRequest (request sent by twitch when the user is redirected)
func (oauth2 *OAuth2) RequestToken(twRequest *http.Request) (TokenResponse, error) {
	authorizationCode := twRequest.URL.Query().Get("code")
	tokenResponse := TokenResponse{}

	if len(authorizationCode) == 0 {
		err := errors.New("The query 'code' parameter is missing, please try again or contact us at contact@wonderstream.tv")
		if twError := twRequest.URL.Query().Get("error"); len(twError) > 0 {
			err = errors.New(twRequest.URL.Query().Get("error_description"))
		}
		return tokenResponse, err
	}

	values := map[string]string{
		"client_id":     oauth2.ClientID,
		"client_secret": oauth2.ClientSecret,
		"code":          authorizationCode,
		"grant_type":    "authorization_code",
		"redirect_uri":  oauth2.RedirectURL,
	}

	jsonRaw, _ := json.Marshal(values)
	resp, err := http.Post("https://api.twitch.tv/api/oauth2/token", "application/json", bytes.NewBuffer(jsonRaw))
	if err != nil {
		return tokenResponse, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return tokenResponse, err
	}

	if len(tokenResponse.AccessToken) <= 0 {
		return tokenResponse, errors.New("Token empty, please ask for a new authorization code")
	}

	return tokenResponse, nil
}
