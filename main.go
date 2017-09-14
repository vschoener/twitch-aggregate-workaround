package main

import (
	"twitch/core"
	"twitch/credential"
	"twitch/storage"
	"twitch/webserver"
)

func main() {
	credential := credential.NewCredential()
	credential.LoadSettings("./parameters.yml")

	database := storage.NewDatabase()
	database.Connect(credential.GetDB())

	defer database.DB.Close()

	oauth2 := core.NewOAuth2(credential.GetTwitch())

	webServer := webserver.NewServer()
	webServer.Start(database, oauth2)
}
