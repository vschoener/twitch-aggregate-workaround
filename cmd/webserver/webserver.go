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
	database.Logger = l.Share()
	database.Logger.SetPrefix("STORAGE")
	database.Connect(credential.GetDB())

	defer database.DB.Close()
	defer l.Close()

	oauth2 := core.NewOAuth2(credential.GetTwitch())
	oauth2Logger := l.Share()
	oauth2Logger.SetPrefix("LIBRARY")
	oauth2.SetLogger(oauth2Logger)

	webServer := webserver.NewServer(credential.ServerSetting)
	webServer.Logger = l.Share()
	webServer.Logger.SetPrefix("WEBSERVER")
	webServer.Start(database, oauth2)
}
