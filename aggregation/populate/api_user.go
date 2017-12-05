package populate

import (
	"github.com/wonderstream/twitch/aggregation/service"
	"github.com/wonderstream/twitch/core"
	twService "github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// User manager
type User struct {
	Loader         *service.Loader
	userRepository repository.UserRepository
	userService    twService.UserService
}

// Initialize channel aggregator
func (u *User) Initialize(loader *service.Loader) {
	u.Loader = loader
	u.userRepository = repository.NewUserRepository(loader.DatabaseManager.Get(storage.DBAggregation), loader.Logger)
	u.userService = twService.NewUserService()
}

// UpdateUser information
func (u *User) UpdateUser(user string, r *core.Request) model.User {
	result := u.userService.GetByChanelNames([]string{user}, r)

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
func (u User) Process(user model.User, token core.TokenResponse) {
	twitchRequest := core.NewRequest(u.Loader.OAuth2, &token)
	twitchRequest.Logger = u.Loader.Logger.Share()
	twitchRequest.Logger.SetPrefix("LIBRARY")

	u.UpdateUser(user.Name, twitchRequest)
}

// End aggregate user
func (u User) End() {

}
