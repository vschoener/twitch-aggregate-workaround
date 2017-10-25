package aggregation

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

// Aggregator interface
type Aggregator interface {
	Initialize(Aggregation)
	Process(model.User, bool, core.TokenResponse)
	End()
}

// Aggregation manager
type Aggregation struct {
	Aggregators     []Aggregator
	OAuth2          *core.OAuth2
	Database        *storage.Database
	Logger          logger.Logger
	twPublicRequest *core.Request
	AppToken        core.TokenResponse
}

// NewAggregation constructor
func NewAggregation(o *core.OAuth2, db *storage.Database, l logger.Logger, appToken core.TokenResponse) Aggregation {
	return Aggregation{
		OAuth2:   o,
		Database: db,
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
}

// Start aggregation process
func (a Aggregation) Start() {
	a.prepare()
	credentialRepository := repository.NewCredentialRepository(a.Database, a.Logger)
	userRepository := repository.NewUserRepository(a.Database, a.Logger)
	users := userRepository.GetUsers()

	for _, user := range users {
		credential, found := credentialRepository.GetCredential(user.Name)
		token := a.AppToken
		if true == found {
			token = transformer.TransformStorageCredentialToCoreTokenResponse(credential)
		}
		for _, aggregator := range a.Aggregators {
			aggregator.Initialize(a)
			aggregator.Process(user, found, token)
			aggregator.End()
		}
	}
}
