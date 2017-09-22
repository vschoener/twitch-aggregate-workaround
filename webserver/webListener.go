package webserver

import (
	"log"
	"net/http"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
)

// Server is use to listen as a webserver
type Server struct {
	ServerSetting
	router Router
}

// ServerSetting will store parameters
type ServerSetting struct {
	Domain string `yaml:"domain"`
	Port   string `yaml:"port"`
}

// NewServer is used to get an instance of WebServer
func NewServer(st ServerSetting) *Server {
	return &Server{
		ServerSetting: st,
	}
}

func (s Server) getAddress() (address string) {
	return s.Domain + ":" + s.Port
}

// Start will listen on the {localURL}:{port}
func (s Server) Start(db *storage.Database, oauth2 *core.OAuth2) {
	// Add require object
	s.router.db = db
	s.router.oauth2 = oauth2

	s.router.Load()
	log.Print("Running web server on " + s.getAddress())
	log.Fatal(http.ListenAndServe(s.getAddress(), nil))
}
