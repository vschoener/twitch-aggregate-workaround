package main

import (
	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
)

func startAggregation(context aggregation.Context) {

	channel := aggregation.Channel{
		Context: context,
	}

	channel.Process()

	// streams := aggregation.Streams{
	// 	Context: context,
	// }
	// streams.Process()

	// users := aggregation.Users{
	// 	Context: context,
	// }
	//
	// users.Process()
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
	credentials := database.GetCredentials()

	twitchRequest := core.NewRequest(oauth2)
	twitchRequest.Logger = l

	context := aggregation.Context{
		OAuth2:      oauth2,
		DB:          database,
		Credentials: credentials,
		Loggger:     l,
		Request:     twitchRequest,
	}

	l.Log("Starting aggregation")
	startAggregation(context)
}
