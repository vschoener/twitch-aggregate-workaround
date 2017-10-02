package webserver

import (
	"log"
	"net/http"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
)

// Server is use to listen as a webserver
type Server struct {
	ServerSetting
	router Router
	Logger logger.Logger
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
	serverRunning := "Running web server on " + s.getAddress()
	log.Print(serverRunning)
	s.Logger.Log(serverRunning)
	err := http.ListenAndServe(s.getAddress(), nil)
	s.Logger.LogInterface(err)
	log.Fatal(err)
}
