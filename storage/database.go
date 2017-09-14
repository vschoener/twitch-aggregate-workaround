package storage

import (
	"database/sql"
	"fmt"
	"log"
	"twitch/core"

	// Only used to query database
	_ "github.com/go-sql-driver/mysql"
)

// Database storage
type Database struct {
	DB *sql.DB
}

// DatabaseSettings info
type DatabaseSettings struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	URL      string `yaml:"url"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

// NewDatabase load the object with credential
func NewDatabase() *Database {
	return &Database{}
}

// Connect to the server
func (s *Database) Connect(dbSettings *DatabaseSettings) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			dbSettings.User,
			dbSettings.Password,
			dbSettings.URL,
			dbSettings.Port,
			dbSettings.Name,
		),
	)
	if err != nil {
		panic(err.Error())
	}
	s.DB = db

	log.Println("Connected to " + dbSettings.URL)
}

// IsConnected return current status
func (s *Database) IsConnected() bool {
	return s.DB != nil
}

// RecordToken used to save token information inside the database
func (s *Database) RecordToken(cs core.ChannelSummary, token core.TokenResponse) {
	stmt, err := s.DB.Prepare("INSERT INTO credential(channelName, access_token, refresh_token, scope, expires_in) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(cs.Name, token.AccessToken, token.RefreshToken, token.Scope, token.ExpiresIn)

	if err != nil {
		log.Fatal(err)
	}
}
