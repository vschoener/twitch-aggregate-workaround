package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wonderstream/twitch/aggregation"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/service"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/repository"
	"github.com/wonderstream/twitch/storage/transformer"
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
		r.logger.Log(fmt.Sprintf("[Router][Request] %#v", req))
		httpHandler.ServeHTTP(w, req)
		r.logger.Log(fmt.Sprintf("[Router][Response] %#v", w))
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
		"addChannel": Route{
			URI:     "/addChannel",
			Handler: r.addChannel,
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

func (r Router) addChannel(w http.ResponseWriter, req *http.Request) {
	s := struct {
		Error       string
		Success     string
		ChannelName string
	}{}

	if req.Method == http.MethodPost {
		req.ParseForm()
		s.ChannelName = req.Form.Get("channel")
		if len(s.ChannelName) > 0 {
			twitchRequest := core.NewRequest(r.oauth2)
			twitchRequest.Logger = r.logger.Share()
			twitchRequest.Logger.SetPrefix("LIBRARY")
			userRepository := repository.NewUserRepository(r.db, r.logger)
			userService := service.NewUserService()
			user := userService.GetByName(s.ChannelName, twitchRequest)
			state := userRepository.StoreUser(transformer.TransformCoreUserToStorageUser(user))

			if state {
				s.Success = "User has been added / updated"
				s.ChannelName = ""
			} else {
				s.Error = "User has not beend added"
			}
		}
	}

	template := core.GetTemplate(core.AddChannelTemplate)
	template.Execute(w, s)
}

// responseOAuthTwitch
func (r Router) responseOAuthTwitch(w http.ResponseWriter, req *http.Request) {
	authAggregation := aggregation.NewAuthAggregation(r.oauth2, r.db)
	err := authAggregation.HandleUserAccessTokenHTTPRequest(w, req, r.logger)

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
