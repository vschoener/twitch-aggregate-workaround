package users

import (
	"fmt"
	"strings"
	"time"

	"github.com/wonderstream/twitch/core"
)

const (
	// UsersURI API ressource
	UsersURI = "/users"
)

// Users manager
type Users struct {
	*core.Request
}

// User information
type User struct {
	DisplayName string    `json:"display_name"`
	ID          int64     `json:"_id,string"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Logo        string    `json:"logo"`
}

// GetByChanelNamesResult contains the result of the request from
// GetByChanelNames()
type GetByChanelNamesResult struct {
	Total int    `json:"_total"`
	Users []User `json:"users"`
}

// GetByChanelNames return a channels list information
func (u Users) GetByChanelNames(channelNames []string) GetByChanelNamesResult {
	result := GetByChanelNamesResult{}
	u.SendRequest(fmt.Sprintf("%s?login=%s", UsersURI, strings.Join(channelNames, ",")), &result)
	return result
}
