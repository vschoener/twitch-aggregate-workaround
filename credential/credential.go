package credential

import (
	"io/ioutil"
	"log"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"

	"gopkg.in/yaml.v2"
)

// Credential manager
type Credential struct {
	core.TwitchSettings      `yaml:"twitch"`
	storage.DatabaseSettings `yaml:"database"`
}

// NewCredential constructor
func NewCredential() *Credential {
	return &Credential{}
}

// LoadSettings from .yml file
func (c *Credential) LoadSettings(path string) *Credential {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if err = yaml.Unmarshal(raw, &c); err != nil {
		panic(err)
	}

	return c
}

// GetTwitch settings
func (c *Credential) GetTwitch() *core.TwitchSettings {
	return &c.TwitchSettings
}

// GetDB settings
func (c *Credential) GetDB() *storage.DatabaseSettings {
	return &c.DatabaseSettings
}
