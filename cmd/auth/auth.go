package main

import (
	"log"

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

	database := storage.NewDatabase()
	database.Logger = l
	dbSetting := c.GetDB(storage.DBAggregation)
	database.Connect(&dbSetting)
	defer database.DB.Close()

	oauth2 := core.NewOAuth2(c.GetTwitch())
	oauth2.Logger = l
	token, err := oauth2.RequestAppAccessToken()

	if err != nil {
		l.LogInterface(err)
		return
	}

	r := repository.CredentialRepository{
		Repository: repository.NewRepository(database, l),
	}
	sToken := transformer.TransformCoreTokenResponseToStorageCredential(token)
	r.StoreCredential(sToken)
}
