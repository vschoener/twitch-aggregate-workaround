package main

import (
	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
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

	r := repository.CredentialRepository{
		Repository: repository.NewRepository(database, l),
	}
	appToken, succeed := r.GetAppToken(oauth2.AppName)

	if !succeed {
		l.LogErrInterface("Credential not found or not loaded properly")
		return
	}

	token := transformer.TransformStorageCredentialToCoreTokenResponse(appToken)

	aggregation := aggregation.NewAggregation(oauth2, database, l, token)

	aggregation.Start()
}
