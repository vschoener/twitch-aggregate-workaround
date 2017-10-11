package aggregation

import (
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

	commonRepository := repository.NewRepository(a.Database, a.Logger)
	u.userRepository = repository.UserRepository{
		Repository: commonRepository,
	}

	u.userService = service.NewUserService()
}

// ProcessUsers aggregate users info
func (u User) processUsers(cr model.Credential) {
	result := u.userService.GetByChanelNames([]string{cr.ChannelName}, u.a.twPublicRequest)

	if result.Total > 0 {
		for _, user := range result.Users {
			u.userRepository.StoreUsers(transformer.TransformCoreUserToStorageUser(user))
		}
	}
}

// Process all Users aggregation
func (u User) Process(cr model.Credential) {
	u.processUsers(cr)
}

// End aggregate user
func (u User) End() {

}
