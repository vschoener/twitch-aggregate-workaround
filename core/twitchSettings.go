package core

import (
	"errors"
	"strings"
)

// TwitchSettings settings
type TwitchSettings struct {
	ClientID              string `yaml:"clientId"`
	ClientSecret          string `yaml:"clientSecret"`
	RedirectURL           string `yaml:"redirectURL"`
	TwitchRequestSettings `yaml:"request"`
	Scopes                []string `ymal:"scopes"`

	// Extra settings for server application
	ErrorRedirectURL   string `yaml:"errorRedirectURL"`
	SuccessRedirectURL string `yaml:"successRedirectURL"`
}

// Check Settings integrity
func (ts TwitchSettings) Check() error {
	var err error
	if len(ts.ClientID) == 0 {
		err = errors.New("ClientID is required")
	} else if len(ts.ClientSecret) == 0 {
		err = errors.New("ClientSecret is required")
	} else if len(ts.RedirectURL) == 0 {
		err = errors.New("Redirect URL is required")
	} else if len(ts.TwitchRequestSettings.URL) == 0 {
		err = errors.New("Twitch Request URL is required")
	} else if len(ts.Scopes) == 0 {
		err = errors.New("Scopes are required")
	}

	return err
}

// TwitchRequestSettings settings
type TwitchRequestSettings struct {
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers,omitempty"`
}

// FormatScopes return the proper scope format
func (ts TwitchSettings) FormatScopes() string {
	return strings.Join(ts.Scopes, " ")
}

const (
	// ChannelURI returns channel uri ressource
	ChannelURI = "/channel"

	// ChannelsURI returns channels uri ressource
	ChannelsURI = "/channels"

	// StreamsURI returns streams uri ressource
	StreamsURI = "/streams"

	// UsersURI returns streams uri ressource
	UsersURI = "/users"
)
