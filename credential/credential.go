package credential

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/webserver"
)

// Interface contains contract
type Interface interface {
	LoadSetting() error
	GetTwitch() *core.TwitchSettings
	GetDB() *storage.DatabaseSettings
	GetLog() *logger.Settings
}

// Credential manager
type Credential struct {
	Interface
	Loader Loader
	Path   string
	AppSetting
}

// DBName type to enum and restrict db name from settings
type DBName string

const (
	// DBAggregation name
	DBAggregation DBName = "aggregation"
	// DBActivity name
	DBActivity DBName = "activity"
)

// AppSetting contains app parameters
type AppSetting struct {
	core.TwitchSettings     `yaml:"twitch"`
	Databases               map[DBName]storage.DatabaseSettings `yaml:"databases"`
	webserver.ServerSetting `yaml:"webserver"`
	LoggerSettings          logger.Settings `yaml:"log"`
}

// NewCredential constructor
func NewCredential(loader Loader, path string) *Credential {
	return &Credential{
		Loader: loader,
		Path:   path,
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
func (c *Credential) GetDB(name DBName) storage.DatabaseSettings {
	return c.Databases[name]
}

// GetLog settings
func (c *Credential) GetLog() *logger.Settings {
	return &c.LoggerSettings
}
