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
	credential := credential.NewCredential()
	credential.LoadSettings("./parameters.yml")

	database := storage.NewDatabase()
	database.Connect(credential.GetDB())

	defer database.DB.Close()

	oauth2 := core.NewOAuth2(credential.GetTwitch())
	credentials := database.GetCredentials()
	context := aggregation.Context{
		OAuth2:      oauth2,
		DB:          database,
		Credentials: credentials,
	}

	startAggregation(context)
}
