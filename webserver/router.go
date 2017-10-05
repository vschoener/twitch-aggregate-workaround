package webserver

import (
	"log"
	"net/http"

	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
)

// HTTPHandler adapter
type HTTPHandler func(http.ResponseWriter, *http.Request)

var routes map[string]HTTPHandler

// Router used to bind any route to a method
type Router struct {
	db     *storage.Database
	oauth2 *core.OAuth2
	Routes map[string]Route
	Mux    *http.ServeMux
	logger logger.Logger
}

// Route info
type Route struct {
	URI     string
	Handler HTTPHandler
}

// Logger is a request logger
func (r Router) Logger(httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.logger.LogInterface(req)
		httpHandler.ServeHTTP(w, req)
		r.logger.LogInterface(w)
	})
}

// Load the requirements (as routes)
func (r *Router) Load() {

	log.Print("Load Router requirements")
	r.Routes = map[string]Route{
		"home": Route{
			URI:     "/",
			Handler: r.home,
		},
		"responseOAuthTwitch": Route{
			URI:     "/responseOAuthTwitch",
			Handler: r.responseOAuthTwitch,
		},
		"connectToTwitch": Route{
			URI:     "/connectToTwitch",
			Handler: r.connectToTwitch,
		},
	}

	r.Mux = http.NewServeMux()

	for _, route := range r.Routes {
		r.Mux.Handle(route.URI, r.Logger(http.HandlerFunc(route.Handler)))
	}

	r.Mux.Handle("/static/", http.FileServer(http.Dir("public")))
}

// home page
func (r Router) home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}

// responseOAuthTwitch
func (r Router) responseOAuthTwitch(w http.ResponseWriter, req *http.Request) {
	authAggregation := aggregation.NewAuthAggregation(r.oauth2, r.db)
	err := authAggregation.HandleHTTPRequest(w, req, r.logger)

	if err != nil {
		http.Redirect(w, req, r.oauth2.ErrorRedirectURL+"?error="+err.Error(), http.StatusSeeOther)
		return
	}
	http.Redirect(w, req, r.oauth2.SuccessRedirectURL, http.StatusSeeOther)
}

// connectToTwitch method to display page with connect button
func (r Router) connectToTwitch(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		template := core.GetTemplate(core.AuthFormTemplate)
		s := struct {
			Error    string
			Settings *core.TwitchSettings
		}{
			Error:    req.URL.Query().Get("error"),
			Settings: r.oauth2.TwitchSettings,
		}
		template.Execute(w, s)
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
