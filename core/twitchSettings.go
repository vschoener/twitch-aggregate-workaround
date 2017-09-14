package core

// TwitchSettings settings
type TwitchSettings struct {
	ClientID              string `yaml:"clientId"`
	ClientSecret          string `yaml:"clientSecret"`
	RedirectURL           string `yaml:"redirectURL"`
	TwitchRequestSettings `yaml:"request"`
}

// TwitchRequestSettings settings
type TwitchRequestSettings struct {
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers,omitempty"`
}
