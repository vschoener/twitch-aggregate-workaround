package aggregation

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// User manager
type User struct {
	Aggregator
	a              Aggregation
	userRepository repository.UserRepository
	userService    service.UserService
}

// Initialize channel aggregator
func (u *User) Initialize(a Aggregation) {
	u.a = a
	u.userRepository = repository.NewUserRepository(a.Database, a.Logger)
	u.userService = service.NewUserService()
}

// UpdateUser information
func (u *User) UpdateUser(user string) model.User {
	result := u.userService.GetByChanelNames([]string{user}, u.a.twPublicRequest)

	model := model.User{}
	if result.Total > 0 {
		for _, user := range result.Users {
			model = transformer.TransformCoreUserToStorageUser(user)
			u.userRepository.StoreUser(model)
		}
	}

	return model
}

// Process all Users aggregation
func (u User) Process(user model.User, isAuthenticated bool, token core.TokenResponse) {
	u.UpdateUser(user.Name)
}

// End aggregate user
func (u User) End() {

}
