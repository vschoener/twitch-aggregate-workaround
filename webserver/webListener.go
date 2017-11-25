package webserver

import (
	"errors"
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

// Check settings integrity
func (st ServerSetting) Check() error {
	var err error
	if len(st.Domain) == 0 {
		err = errors.New("Domain is required")
	} else if len(st.Port) == 0 {
		err = errors.New("port is required")
	}

	return err
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
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
	s.router.logger = s.Logger

	s.router.Load()
	s.Logger.Log("Running web server on " + s.getAddress())
	err := http.ListenAndServe(s.getAddress(), s.router.Mux)
	s.Logger.LogInterface(err)
	log.Fatal(err)
}
