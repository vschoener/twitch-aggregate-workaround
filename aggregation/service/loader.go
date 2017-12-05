package service

import (
	"log"
	"os"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/credential"
	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
)

// LoaderInterface is an interface of Loader
type LoaderInterface interface {
	Initialize()
	HadError()
	GetError()
}

// Service type to handle type service
type Service int

const (
	// OAUTH2 service
	OAUTH2 Service = iota
	// Logger service
	Logger
	// DatabaseManager service
	DatabaseManager
	// Credential service
	Credential
)

// Loader handles basic aggregations Initializing
type Loader struct {
	Credential      *credential.Credential
	OAuth2          *core.OAuth2
	Logger          logger.Logger
	DatabaseManager *storage.DatabaseManager
	err             error
}

// Initialize aggregation requirements
func (l *Loader) Initialize() {
	if l.initializeParameters(); l.err != nil {
		log.Fatal(l.err)
		os.Exit(-1)
	}

	l.initialzeLogger()
	l.initializeDatabases()
	l.initializeOAuth2Token()
}

func (l *Loader) initializeParameters() *Loader {
	l.Credential = credential.NewCredential(credential.YAMLLoader{}, "./parameters.yml")
	if err := l.Credential.LoadSetting(); err != nil {
		l.err = err
	}

	return l
}

// initializeDatabases
func (l *Loader) initializeDatabases() *Loader {
	l.DatabaseManager = storage.GetDM()
	l.DatabaseManager.ConnectNewDatabase(storage.DBAggregation, l.Credential.GetDB(storage.DBAggregation), l.Logger.Share())
	l.DatabaseManager.ConnectNewDatabase(storage.DBActivity, l.Credential.GetDB(storage.DBActivity), l.Logger.Share())

	return l
}

// initialzeLogger
func (l *Loader) initialzeLogger() *Loader {
	l.Logger = logger.NewLogger()
	l.Logger.Connect(l.Credential.LoggerSettings)

	return l
}

// Close should be use to close deferred function (any service here will be close)
func (l *Loader) Close() *Loader {
	l.Logger.Close()
	l.DatabaseManager.Get(storage.DBAggregation).Close()
	l.DatabaseManager.Get(storage.DBActivity).Close()

	return l
}

// initialzeLogger
func (l *Loader) initializeOAuth2Token() {
	l.OAuth2 = core.NewOAuth2(l.Credential.GetTwitch())
	oauth2Logger := l.Logger.Share()
	oauth2Logger.SetPrefix("LIBRARY")
	l.OAuth2.SetLogger(oauth2Logger)
}

// HadError let us if anything is wrong during the process
func (l *Loader) HadError() bool {
	return l.err != nil
}

// GetError returns error
func (l *Loader) GetError() error {
	return l.err
}
