package aggregation

import (
	"fmt"
	"reflect"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// Aggregator interface
type Aggregator interface {
	Initialize(*Aggregation)
	Process(model.User, bool, core.TokenResponse)
	End()
}

// Aggregation manager
type Aggregation struct {
	Aggregators     []Aggregator
	OAuth2          *core.OAuth2
	DM              *storage.DatabaseManager
	Logger          logger.Logger
	twPublicRequest *core.Request
	AppToken        core.TokenResponse
}

// NewAggregation constructor
func NewAggregation(o *core.OAuth2, dm *storage.DatabaseManager, l logger.Logger, appToken core.TokenResponse) Aggregation {
	return Aggregation{
		OAuth2:   o,
		DM:       dm,
		Logger:   l,
		AppToken: appToken,
	}
}

func (a *Aggregation) prepare() {
	// Prepare Non Auth request to avoid building the same again and again
	twitchRequest := core.NewRequest(a.OAuth2)
	twitchRequest.Logger = a.Logger.Share()
	twitchRequest.Logger.SetPrefix("LIBRARY")
	a.twPublicRequest = twitchRequest

	a.Aggregators = append(a.Aggregators, &Channel{})
	a.Aggregators = append(a.Aggregators, &User{})
	a.Aggregators = append(a.Aggregators, &Stream{})
	a.Aggregators = append(a.Aggregators, &Summarize{})
}

// Start aggregation process
func (a Aggregation) Start() {
	a.prepare()
	credentialRepository := repository.NewCredentialRepository(a.DM.Get(storage.DBAggregation), a.Logger)
	userRepository := repository.NewUserRepository(a.DM.Get(storage.DBAggregation), a.Logger)
	users := userRepository.GetUsers()

	for _, user := range users {
		credential, found := credentialRepository.GetCredential(user.Name)
		token := a.AppToken
		if true == found {
			token = transformer.TransformStorageCredentialToCoreTokenResponse(credential)
		}
		for _, aggregator := range a.Aggregators {
			a.Logger.Log(fmt.Sprintf("Aggregation %s started for %s", reflect.TypeOf(aggregator).String(), user.Name))
			aggregator.Initialize(&a)
			aggregator.Process(user, found, token)
			aggregator.End()
		}
	}
}
