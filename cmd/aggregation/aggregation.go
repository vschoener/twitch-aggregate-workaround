package main

import (
	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/storage"
)

func startAggregation(context aggregation.Context) {
	channel := aggregation.Channel{
		Context: context,
	}

	channel.Process()
}

func main() {
	credential := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	credential.LoadSetting()

	database := storage.NewDatabase()
	database.Connect(credential.GetDB())

	defer database.DB.Close()

	oauth2 := core.NewOAuth2(credential.GetTwitch())

	// TODO: It's not required not but we should itenarate by range later
	credentials := database.GetCredentials()
	context := aggregation.Context{
		OAuth2:      oauth2,
		DB:          database,
		Credentials: credentials,
	}

	startAggregation(context)
}
