package aggregation

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
	"github.com/wonderstream/twitch/storage/repository"
)

// Aggregator interface
type Aggregator interface {
	Initialize(Aggregation)
	Process(model.Credential)
	End()
}

// Aggregation manager
type Aggregation struct {
	Aggregators     []Aggregator
	OAuth2          *core.OAuth2
	Database        *storage.Database
	Logger          logger.Logger
	twPublicRequest *core.Request
}

// NewAggregation constructor
func NewAggregation(o *core.OAuth2, db *storage.Database, l logger.Logger) Aggregation {
	return Aggregation{
		OAuth2:   o,
		Database: db,
		Logger:   l,
	}
}

func (a *Aggregation) prepare() {
	// Prepare Non Auth request to avoid building the same again and again
	twitchRequest := core.NewRequest(a.OAuth2)
	twitchRequest.Logger = a.Logger
	a.twPublicRequest = twitchRequest

	a.Aggregators = append(a.Aggregators, &Channel{})
	a.Aggregators = append(a.Aggregators, &User{})
	a.Aggregators = append(a.Aggregators, &Stream{})
}

// Start aggregation process
func (a Aggregation) Start() {
	a.prepare()

	credentialRepository := repository.CredentialRepository{
		Repository: repository.NewRepository(a.Database, a.Logger),
	}
	credentials := credentialRepository.GetCredentials()

	for _, aggregator := range a.Aggregators {
		aggregator.Initialize(a)
		for _, credential := range credentials {
			aggregator.Process(credential)
		}
		aggregator.End()
	}
}
