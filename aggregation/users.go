package aggregation

import (
	"github.com/wonderstream/twitch/core/api/users"
)

// Users manager
type Users struct {
	Context
}

// ProcessUsers aggregate users info
func (u Users) ProcessUsers() {

	users := users.Users{
		Request: u.Request,
	}

	result := users.GetByChanelNames([]string{"elynia_gaming"})

	if result.Total > 0 {
		for _, user := range result.Users {
			u.Context.DB.StoreUsers(user)
		}
	}
}

// Process all Users aggregation
func (u Users) Process() {
	u.ProcessUsers()
}
