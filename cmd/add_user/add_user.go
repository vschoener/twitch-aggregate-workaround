package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
)

func help() {
	fmt.Println("Help : Use the following command to add a user")
	fmt.Println(">./add_user [user1, orChannel1, channel2, channel3]")
}

func main() {

	args := os.Args[1:]

	if (len(args) == 0) || (args[0] == "help") {
		help()
		return
	}

	c := credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	if err := c.LoadSetting(); err != nil {
		log.Fatal(err)
	}

	l := logger.NewLogger()
	l.Connect(c.LoggerSettings)
	defer l.Close()

	dm := storage.GetDM()
	dm.ConnectNewDatabase(storage.DBAggregation, c.GetDB(storage.DBAggregation), l.Share())
	defer dm.Get(storage.DBAggregation).Close()

	oauth2 := core.NewOAuth2(c.GetTwitch())
	oauth2Logger := l.Share()
	oauth2Logger.SetPrefix("LIBRARY")
	oauth2.SetLogger(oauth2Logger)

	twitchRequest := core.NewRequest(oauth2, nil)
	twitchRequest.Logger = l.Share()
	twitchRequest.Logger.SetPrefix("LIBRARY")

	userRepository := repository.NewUserRepository(dm.Get(storage.DBAggregation), l)
	userService := service.NewUserService()

	for _, user := range args {
		twUser, err := userService.GetByName(user, twitchRequest)
		if nil == err {
			l.Log("User found")
			state := userRepository.StoreUser(transformer.TransformCoreUserToStorageUser(twUser))
			if state {
				l.Log("User added")
			} else {
				l.Log("User not added")
			}
		} else {
			l.Log("User not found... Aborted")
		}
	}
}
