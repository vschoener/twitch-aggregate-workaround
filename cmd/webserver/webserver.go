package main

import (
	"log"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/webserver"
)

func main() {
	log.Println("Preparing and loading webserver requirements...")
	credential := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	credential.LoadSetting()

	l := logger.NewLogger()
	l.Connect(credential.LoggerSettings)

	database := storage.NewDatabase()
	database.Logger = l
	database.Connect(credential.GetDB())

	defer database.DB.Close()
	defer (l.(*logger.LogEntry)).Client.Close()

	oauth2 := core.NewOAuth2(credential.GetTwitch())
	oauth2.SetLogger(l)

	webServer := webserver.NewServer(credential.ServerSetting)
	webServer.Logger = l
	webServer.Start(database, oauth2)
}
