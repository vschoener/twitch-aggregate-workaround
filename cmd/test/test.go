package main

import (
	"log"
	"time"

	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/repository"
)

func main() {
	c := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	c.LoadSetting()

	l := logger.NewLogger()
	l.Connect(c.LoggerSettings)
	defer l.Close()

	l.Log("Preparing aggregation")

	dm := storage.GetDM()
	dm.ConnectNewDatabase(storage.DBAggregation, c.GetDB(storage.DBAggregation), l.Share())
	defer dm.Get(storage.DBAggregation).Close()
	dm.ConnectNewDatabase(storage.DBActivity, c.GetDB(storage.DBActivity), l.Share())
	defer dm.Get(storage.DBActivity).Close()

	//oauth2 := core.NewOAuth2(c.GetTwitch())
	// twitchRequest := core.NewRequest(oauth2)
	// twitchRequest.Logger = l
	// v := service.NewVideoService()
	// log.Printf("%#v", v.GetVideosFromID(42544623, twitchRequest, 10))

	r := repository.NewActivityStorageRepository(dm.Get(storage.DBActivity), l)
	dateStart := time.Date(2017, time.November, 20, 0, 0, 0, 0, time.UTC)
	filters := storage.QueryFilter{
		DateStart: &dateStart,
		Exclude: map[string][]string{
			"username": {"wondrlurker"},
		},
		Include: map[string][]string{
			"type": {"JOIN", "PART"},
		},
	}

	//activities := r.GetChannelActivities("dota2ti", filters)
	hoursWatched := r.GetWatchedTime("dota2ti", filters)
	log.Printf("%d", hoursWatched)
}
