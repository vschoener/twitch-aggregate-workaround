package core

import "strings"

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

// TwitchRequestSettings settings
type TwitchRequestSettings struct {
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers,omitempty"`
}

// FormatScopes return the proper scope format
func (ts TwitchSettings) FormatScopes() string {
	return strings.Join(ts.Scopes, " ")
}
