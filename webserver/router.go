package webserver

import (
	"log"
	"net/http"
	"twitch/core"
	"twitch/storage"
)

// Router used to bind any route to a method
type Router struct {
	db     *storage.Database
	oauth2 *core.OAuth2
}

// Load the requirements (as routes)
func (r Router) Load() {
	log.Print("Load Router requirements")
	http.HandleFunc("/", r.home)
	http.HandleFunc("/connectToTwitch", r.connectToTwitch)
	http.HandleFunc("/responseOAuthTwitch", r.responseOAuthTwitch)
	http.HandleFunc("/getChannelInfo", r.getChannelInfo)
	http.Handle("/static/", http.FileServer(http.Dir("public")))
}

// home page
func (r Router) home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}

// responseOAuthTwitch
func (r Router) responseOAuthTwitch(w http.ResponseWriter, req *http.Request) {
	token, err := r.oauth2.RequestToken(req)
	if err != nil {
		http.Redirect(w, req, "http://localhost:8080/connectToTwitch", http.StatusUnauthorized)
		return
	}

	twitchRequest := core.NewRequest(r.oauth2, token)
	channel := core.Channel{Request: twitchRequest}
	channelSummary := channel.RequestSummary()

	r.db.RecordToken(channelSummary, token)
}

// getChannelInfo
func (r Router) getChannelInfo(w http.ResponseWriter, req *http.Request) {

}

// connectToTwitch method to display page with connect button
func (r Router) connectToTwitch(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		template := core.GetTemplate(core.AuthFormTemplate)
		template.Execute(w, r.oauth2.TwitchSettings)
	} else {
		w.Write([]byte("This method is not handled"))
	}
}
