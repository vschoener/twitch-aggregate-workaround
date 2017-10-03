package storage

import (
	"database/sql"
	"fmt"

	// Only used to query database
	_ "github.com/go-sql-driver/mysql"
	"github.com/wonderstream/twitch/logger"
)

// Database storage
type Database struct {
	DB     *sql.DB
	Logger logger.Logger
}

// DatabaseSettings info
type DatabaseSettings struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	URL      string `yaml:"url"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

// QueryLogger used to log or debug query
type QueryLogger struct {
	Query      string
	Parameters map[string]interface{}
}

// NewDatabase load the object with credential
func NewDatabase() *Database {
	return &Database{}
}

// Connect to the server
func (s *Database) Connect(dbSettings *DatabaseSettings) {

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbSettings.User,
			dbSettings.Password,
			dbSettings.URL,
			dbSettings.Port,
			dbSettings.Name,
		),
	)
	if err != nil {
		s.Logger.Log(err.Error())
		panic(err.Error())
	}
	s.DB = db

	s.Logger.Log("Database storage connected to " + dbSettings.URL)
}

// IsConnected return current status
func (s *Database) IsConnected() bool {
	return s.DB != nil
}
