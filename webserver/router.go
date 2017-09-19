package webserver

import (
	"log"
	"net/http"

	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
)

// Router used to bind any route to a method
type Router struct {
	db     *storage.Database
	oauth2 *core.OAuth2
	Routes *map[string]string
}

// Load the requirements (as routes)
func (r Router) Load() {
	log.Print("Load Router requirements")
	http.HandleFunc("/", r.home)
	http.HandleFunc("/connectToTwitch", r.connectToTwitch)
	http.HandleFunc("/responseOAuthTwitch", r.responseOAuthTwitch)
	http.Handle("/static/", http.FileServer(http.Dir("public")))
}

// home page
func (r Router) home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}

// responseOAuthTwitch
func (r Router) responseOAuthTwitch(w http.ResponseWriter, req *http.Request) {
	authAggregation := aggregation.NewAuthAggregation(r.oauth2, r.db)
	authAggregation.HandleHTTPRequest(w, req)
	http.Redirect(w, req, r.oauth2.SuccessRedirectURL, http.StatusSeeOther)
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

func (r Router) handleError(err error, w http.ResponseWriter, req *http.Request) {
	if err != nil {
		http.Redirect(w, req, r.oauth2.ErrorRedirectURL, http.StatusUnauthorized)
		return
	}
}
