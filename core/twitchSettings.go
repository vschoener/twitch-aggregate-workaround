package core

import "strings"

// TwitchSettings settings
type TwitchSettings struct {
	ClientID              string `yaml:"clientId"`
	ClientSecret          string `yaml:"clientSecret"`
	RedirectURL           string `yaml:"redirectURL"`
	TwitchRequestSettings `yaml:"request"`
	Scopes                []string `ymal:"scopes"`
	AppName               string   `yaml:"app_name"`

	// Extra settings for server application
	ErrorRedirectURL   string `yaml:"errorRedirectURL"`
	SuccessRedirectURL string `yaml:"successRedirectURL"`
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
