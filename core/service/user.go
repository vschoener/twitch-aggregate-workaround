package service

import (
	"fmt"
	"strings"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
)

// UserService handles processes for the user channel
type UserService struct {
	*core.Request
}

// GetByChanelNamesResult contains the result of the request from
// GetByChanelNames()
type GetByChanelNamesResult struct {
	Total int          `json:"_total"`
	Users []model.User `json:"users"`
}

// GetByChanelNames return a channels list information
func (s UserService) GetByChanelNames(channelNames []string) GetByChanelNamesResult {
	result := GetByChanelNamesResult{}
	s.SendRequest(fmt.Sprintf("%s?login=%s", core.UsersURI, strings.Join(channelNames, ",")), &result)

	return result
}
