package main

import (
	"log"

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
	if err := c.LoadSetting(); err != nil {
		log.Fatal(err)
	}

	l := logger.NewLogger()
	l.Connect(c.LoggerSettings)
	defer l.Close()

	l.Log("Preparing aggregation")

	dm := storage.GetDM()
	dm.ConnectNewDatabase(storage.DBAggregation, c.GetDB(storage.DBAggregation), l.Share())
	defer dm.Get(storage.DBAggregation).Close()
	dm.ConnectNewDatabase(storage.DBActivity, c.GetDB(storage.DBActivity), l.Share())
	defer dm.Get(storage.DBActivity).Close()

	oauth2 := core.NewOAuth2(c.GetTwitch())
	oauth2Logger := l.Share()
	oauth2Logger.SetPrefix("LIBRARY")
	oauth2.SetLogger(oauth2Logger)

	r := repository.CredentialRepository{
		Repository: repository.NewRepository(dm.Get(storage.DBAggregation), l),
	}
	appToken, succeed := r.GetAppToken(oauth2.AppName)

	if !succeed {
		l.LogErrInterface("AppToken not found, did you authenticate the script using the auth binary ?")
		return
	}

	token := transformer.TransformStorageCredentialToCoreTokenResponse(appToken)

	aggregationLogger := l.Share()
	aggregationLogger.SetPrefix("AGGREGATION")
	aggregation := aggregation.NewAggregation(
		oauth2,
		dm,
		aggregationLogger,
		token,
	)

	aggregation.Start()
}
