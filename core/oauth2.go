package core

import (
	"strings"
	"time"

	"github.com/wonderstream/twitch/logger"
)

// OAuth2 manager
type OAuth2 struct {
	Scopes map[string]string
	*TwitchSettings
	Logger logger.Logger
}

// TokenResponse use to store response body when getting a new token
type TokenResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int64    `json:"expires_in"`
	Scopes       []string `json:"scope"`
}

// IsTokenValid checks if the token is still valide
func (s TokenResponse) IsTokenValid(date time.Time) bool {
	if s.ExpiresIn == 0 {
		return true
	}

	return date.Add(time.Second * time.Duration(s.ExpiresIn)).Before(time.Now())
}

// FormatScopes return the proper scope format
func (s TokenResponse) FormatScopes() string {
	return strings.Join(s.Scopes, " ")
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

// SetLogger set logger service
func (oauth2 *OAuth2) SetLogger(logger logger.Logger) {
	oauth2.Logger = logger
}

// RequestToken only send an auth request to retrieve an Access token
func (oauth2 *OAuth2) RequestToken(values map[string]string) (TokenResponse, error) {
	tokenResponse := TokenResponse{}
	values["client_id"] = oauth2.ClientID
	values["client_secret"] = oauth2.ClientSecret

	request := NewRequest(oauth2)
	request.Logger = oauth2.Logger
	request.SetPost(values, "application/json")
	err := request.SendRequest("/oauth2/token", &tokenResponse)

	return tokenResponse, err
}

// RequestUserAcccessToken has to request a token using Query "code" paramater available
// inside the twRequest (request sent by twitch when the user is redirected)
func (oauth2 *OAuth2) RequestUserAcccessToken(authorizationCode string) (TokenResponse, error) {
	values := map[string]string{
		"code":         authorizationCode,
		"grant_type":   "authorization_code",
		"redirect_uri": oauth2.RedirectURL,
	}
	tokenResponse, err := oauth2.RequestToken(values)

	return tokenResponse, err
}

// RequestAppAccessToken ask a server to server token
func (oauth2 *OAuth2) RequestAppAccessToken() (TokenResponse, error) {
	values := map[string]string{
		"grant_type": "client_credentials",
		"scope":      oauth2.FormatScopes(),
	}
	tokenResponse, err := oauth2.RequestToken(values)

	return tokenResponse, err
}

// RefreshToken request a new token
func (oauth2 *OAuth2) RefreshToken(token TokenResponse) (TokenResponse, error) {
	values := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": token.RefreshToken,
	}
	token, err := oauth2.RequestToken(values)
	if err != nil {
		return token, err
	}

	return token, nil
}
