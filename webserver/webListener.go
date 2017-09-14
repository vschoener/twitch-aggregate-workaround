package webserver

import (
	"log"
	"net/http"
	"twitch/core"
	"twitch/storage"
)

// Server is use to listen as a webserver
type Server struct {
	localURL string
	port     string
	router   Router
}

// NewServer is used to get an instance of WebServer
func NewServer() *Server {
	return &Server{
		localURL: "localhost",
		port:     "8080",
	}
}

func (s Server) getAddress() (address string) {
	return s.localURL + ":" + s.port
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
