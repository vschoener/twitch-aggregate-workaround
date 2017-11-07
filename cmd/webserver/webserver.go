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
	c := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	c.LoadSetting()

	l := logger.NewLogger()
	l.Connect(c.LoggerSettings)

	database := storage.NewDatabase()
	database.Logger = l.Share()
	database.Logger.SetPrefix("STORAGE")
	dbSetting := c.GetDB(credential.DBAggregation)
	database.Connect(&dbSetting)

	defer database.DB.Close()
	defer l.Close()

	oauth2 := core.NewOAuth2(c.GetTwitch())
	oauth2Logger := l.Share()
	oauth2Logger.SetPrefix("LIBRARY")
	oauth2.SetLogger(oauth2Logger)

	webServer := webserver.NewServer(c.ServerSetting)
	webServer.Logger = l.Share()
	webServer.Logger.SetPrefix("WEBSERVER")
	webServer.Start(database, oauth2)
}
