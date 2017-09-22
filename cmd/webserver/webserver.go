package main

import (
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/webserver"
)

func main() {
	credential := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	credential.LoadSetting()

	database := storage.NewDatabase()
	database.Connect(credential.GetDB())

	defer database.DB.Close()

	oauth2 := core.NewOAuth2(credential.GetTwitch())

	webServer := webserver.NewServer()
	webServer.Start(database, oauth2)
}
