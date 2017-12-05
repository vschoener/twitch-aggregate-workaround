package aggregation

import (
	"fmt"
	"reflect"

	"github.com/wonderstream/twitch/aggregation/service"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// Aggregator interface
type Aggregator interface {
	Initialize(*service.Loader)
	Process(model.User, core.TokenResponse)
	End()
}

// AggregatorManager manager
type AggregatorManager struct {
	Aggregators []Aggregator
}

// Start aggregation process
func (a AggregatorManager) Start(l *service.Loader) {
	credentialRepository := repository.NewCredentialRepository(l.DatabaseManager.Get(storage.DBAggregation), l.Logger)
	userRepository := repository.NewUserRepository(l.DatabaseManager.Get(storage.DBAggregation), l.Logger)
	users := userRepository.GetUsers()

	for _, user := range users {
		credential, found := credentialRepository.GetCredential(user.Name)
		var token core.TokenResponse
		if true == found {
			token = transformer.TransformStorageCredentialToCoreTokenResponse(credential)
		}
		for _, aggregator := range a.Aggregators {
			l.Logger.Log(fmt.Sprintf("Aggregation %s started for %s", reflect.TypeOf(aggregator).String(), user.Name))
			aggregator.Initialize(l)
			aggregator.Process(user, token)
			aggregator.End()
		}
	}
}
