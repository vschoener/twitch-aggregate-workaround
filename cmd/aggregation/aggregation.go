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
	c := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	c.LoadSetting()

	l := logger.NewLogger()
	l.Connect(c.LoggerSettings)
	defer l.Close()

	l.Log("Preparing aggregation")

	database := storage.NewDatabase()
	database.Logger = l.Share()
	database.Logger.SetPrefix("STORAGE")
	dbSetting := c.GetDB(credential.DBAggregation)
	database.Connect(&dbSetting)
	defer database.DB.Close()

	oauth2 := core.NewOAuth2(c.GetTwitch())
	oauth2Logger := l.Share()
	oauth2Logger.SetPrefix("LIBRARY")
	oauth2.SetLogger(oauth2Logger)

	r := repository.CredentialRepository{
		Repository: repository.NewRepository(database, l),
	}
	appToken, succeed := r.GetAppToken(oauth2.AppName)

	if !succeed {
		l.LogErrInterface("Credential not found or not loaded properly")
		return
	}

	token := transformer.TransformStorageCredentialToCoreTokenResponse(appToken)

	aggregationLogger := l.Share()
	aggregationLogger.SetPrefix("AGGREGATION")
	aggregation := aggregation.NewAggregation(
		oauth2,
		database,
		aggregationLogger,
		token,
	)

	aggregation.Start()
}
