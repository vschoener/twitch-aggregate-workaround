package main

import (
	"log"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
)

func main() {
	c := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	c.LoadSetting()

	l := logger.NewLogger()
	l.Connect(c.LoggerSettings)
	defer l.(*logger.LogEntry).Client.Close()

	l.Log("Preparing aggregation")

	database := storage.NewDatabase()
	database.Logger = l
	dbSetting := c.GetDB(storage.DBAggregation)
	database.Connect(&dbSetting)
	defer database.DB.Close()

	oauth2 := core.NewOAuth2(c.GetTwitch())

	twitchRequest := core.NewRequest(oauth2)
	twitchRequest.Logger = l
	v := service.NewVideoService()

	log.Printf("%#v", v.GetVideosFromID(42544623, twitchRequest, 10))
}
