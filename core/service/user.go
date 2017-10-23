package service

import (
	"fmt"
	"strings"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
)

// UserService handles processes for the user channel
type UserService struct {
}

// NewUserService constructor
func NewUserService() UserService {
	return UserService{}
}

// GetByChanelNamesResult contains the result of the request from
// GetByChanelNames()
type GetByChanelNamesResult struct {
	Total int          `json:"_total"`
	Users []model.User `json:"users"`
}

// GetByChanelNames return a channels list information
func (s UserService) GetByChanelNames(channelNames []string, r *core.Request) GetByChanelNamesResult {
	result := GetByChanelNamesResult{}
	r.SendRequest(fmt.Sprintf("%s?login=%s", core.UsersURI, strings.Join(channelNames, ",")), &result)

	return result
}

// GetByName retrieve user information from channel or user name
func (s UserService) GetByName(u string, r *core.Request) model.User {
	result := s.GetByChanelNames([]string{u}, r)

	user := model.User{}
	if result.Total == 1 {
		user = result.Users[0]
	}

	return user
}
