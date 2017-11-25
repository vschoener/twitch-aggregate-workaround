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

// Setting Interface
type Setting interface {
	Check() error
}

// Credential manager
type Credential struct {
	Interface
	Loader Loader
	Path   string
	AppSetting
	checkersList map[SettingName]bool
}

// DBName type to enum and restrict db name from settings
type DBName string

const (
	// DBAggregation name
	DBAggregation DBName = "aggregation"
	// DBActivity name
	DBActivity DBName = "activity"
)

// SettingName type to enum and list settings to check
type SettingName int

const (
	// SettingTW name
	SettingTW SettingName = iota + 0
	// SettingDB name
	SettingDB
	// SettingLog Name
	SettingLog
	// SettingWS Name
	SettingWS
)

// AppSetting contains app parameters
type AppSetting struct {
	TwitchSettings          core.TwitchSettings                 `yaml:"twitch"`
	Databases               map[DBName]storage.DatabaseSettings `yaml:"databases"`
	webserver.ServerSetting `yaml:"webserver"`
	LoggerSettings          logger.Settings `yaml:"log"`
}

// NewCredential constructor
func NewCredential(loader Loader, path string) *Credential {
	return &Credential{
		Loader: loader,
		Path:   path,
		checkersList: map[SettingName]bool{
			SettingTW:  true,
			SettingDB:  true,
			SettingWS:  true,
			SettingLog: true,
		},
	}
}

// SetCheckSetting changes state of the checker
func (c *Credential) SetCheckSetting(s SettingName, state bool) {
	c.checkersList[s] = state
}

// LoadSetting from .yml file
func (c *Credential) LoadSetting() error {
	err := c.Loader.Load(c.Path, &c.AppSetting)

	if err == nil {
		if c.checkersList[SettingTW] {
			if err = c.AppSetting.TwitchSettings.Check(); err != nil {
				return err
			}
		}
		if c.checkersList[SettingDB] {
			for _, database := range c.AppSetting.Databases {
				if err = database.Check(); err != nil {
					return err
				}
			}
		}
		if c.checkersList[SettingWS] {
			if err = c.AppSetting.ServerSetting.Check(); err != nil {
				return err
			}
		}
		if c.checkersList[SettingLog] {
			if err = c.AppSetting.LoggerSettings.Check(); err != nil {
				return err
			}
		}
	}

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
