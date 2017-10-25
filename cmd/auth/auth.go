package main

import (
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
	defer l.Close()

	database := storage.NewDatabase()
	database.Logger = l
	database.Connect(credential.GetDB())
	defer database.DB.Close()

	oauth2 := core.NewOAuth2(credential.GetTwitch())
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
	sToken.AppName = oauth2.AppName
	r.SaveAppCredential(sToken)
}
