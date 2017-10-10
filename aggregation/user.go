package aggregation

import (
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// User manager
type User struct {
	Context
	repository.UserRepository
}

// ProcessUsers aggregate users info
func (u User) ProcessUsers() {

	userService := service.UserService{Request: u.Request}

	result := userService.GetByChanelNames([]string{"elynia_gaming"})

	if result.Total > 0 {
		for _, user := range result.Users {
			u.StoreUsers(transformer.TransformCoreUserToStorageUser(user))
		}
	}
}

// Process all Users aggregation
func (u User) Process() {
	u.ProcessUsers()
}
