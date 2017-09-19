package aggregation

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
)

// Context contains object to request and store the data
type Context struct {
	OAuth2      *core.OAuth2
	DB          *storage.Database
	Credentials []storage.Credential
}
