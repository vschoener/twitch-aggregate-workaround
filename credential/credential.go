package credential

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
)

// Credential manager
type Credential struct {
	Loader Loader
	Path   string
	AppSetting
}

// AppSetting contains app parameters
type AppSetting struct {
	core.TwitchSettings      `yaml:"twitch"`
	storage.DatabaseSettings `yaml:"database"`
}

// NewCredential constructor
func NewCredential(loader Loader, path string) *Credential {
	return &Credential{
		Loader: loader,
		Path:   path,
		AppSetting: AppSetting{
			TwitchSettings:   core.TwitchSettings{},
			DatabaseSettings: storage.DatabaseSettings{},
		},
	}
}

// LoadSetting from .yml file
func (c *Credential) LoadSetting() error {
	err := c.Loader.Load(c.Path, &c.AppSetting)

	return err
}

// GetTwitch settings
func (c *Credential) GetTwitch() *core.TwitchSettings {
	return &c.TwitchSettings
}

// GetDB settings
func (c *Credential) GetDB() *storage.DatabaseSettings {
	return &c.DatabaseSettings
}
