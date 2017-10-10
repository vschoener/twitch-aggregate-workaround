package main

import (
	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/repository"
)

func startAggregation(c aggregation.Context) {

	c.Loggger.Log("Starting aggregation")

	commonRepository := repository.NewRepository(c.DB, c.Loggger)
	channelRepository := repository.ChannelRepository{
		Repository: commonRepository,
	}
	videoRepository := repository.ChannelVideoRepository{
		Repository: commonRepository,
	}

	userRepository := repository.UserRepository{
		Repository: commonRepository,
	}

	// for _, credential := range c.Credentials {
	// 	twitchRequest := core.NewUserAccessTokenRequest(c.OAuth2, credential.TokenResponse)
	// 	twitchRequest.Logger = c.Loggger
	// 	channel := service.ChannelService{Request: twitchRequest}
	// 	channelSummary := channel.GetInfo()
	// 	c.StoreChannelSummary(channelSummary)
	// }

	channel := aggregation.Channel{
		Context:                c,
		ChannelRepository:      channelRepository,
		ChannelVideoRepository: videoRepository,
	}

	channel.Process()

	// streams := aggregation.Streams{
	// 	Context: context,
	// }
	// streams.Process()

	users := aggregation.User{
		Context:        c,
		UserRepository: userRepository,
	}

	users.Process()
}

func main() {

	credential := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	credential.LoadSetting()

	l := logger.NewLogger()
	l.Connect(credential.LoggerSettings)
	defer l.(*logger.LogEntry).Client.Close()

	l.Log("Preparing aggregation")

	database := storage.NewDatabase()
	database.Logger = l
	database.Connect(credential.GetDB())
	defer database.DB.Close()

	oauth2 := core.NewOAuth2(credential.GetTwitch())

	// TODO: It's not required now but we should itenarate by range later
	credentialRepository := repository.CredentialRepository{
		Repository: repository.NewRepository(database, l),
	}
	credentials := credentialRepository.GetCredentials()

	// Prepare Non Auth request to avoid building the same again and again
	twitchRequest := core.NewRequest(oauth2)
	twitchRequest.Logger = l

	context := aggregation.Context{
		OAuth2:      oauth2,
		DB:          database,
		Credentials: credentials,
		Loggger:     l,
		Request:     twitchRequest,
	}

	startAggregation(context)
}
