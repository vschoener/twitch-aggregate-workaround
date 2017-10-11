package main

import (
	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
)

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

	aggregation := aggregation.NewAggregation(oauth2, database, l)
	aggregation.Start()
}
