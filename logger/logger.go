package logger

import (
	"fmt"
	"log"

	"github.com/bsphere/le_go"
)

// Logger interface
type Logger interface {
	// This method should load the source define as string to store element
	// in the interface
	Log(string) error

	// LogInterface logg any complex struct
	LogInterface(interface{}) error
	// LogErrInterface shortcut to tag Error
	LogErrInterface(interface{}) error

	Connect(Settings)
}

// LogEntry contains requirements to log
type LogEntry struct {
	Logger
	Client *le_go.Logger
	Settings
}

// Settings contains credential
type Settings struct {
	State   bool   `yaml:"state"`
	Key     string `yaml:"key"`
	Verbose bool   `yaml:"verbose"`
}

// NewLogger constructor
func NewLogger() Logger {
	return &LogEntry{}
}

// Connect connect to LogEntries server
func (le *LogEntry) Connect(s Settings) {

	le.Settings = s
	if len(s.Key) == 0 {
		le.Client = nil
		return
	}

	var err error
	le.Client, err = le_go.Connect(s.Key)

	if err != nil {
		log.Fatal(err)
	}

	le.Log("Logger is connected to LogEntries using key " + s.Key)
}

// Log simple string
func (le *LogEntry) Log(message string) error {

	if le.Verbose {
		log.Println(message)
	}

	if false == le.State {
		return nil
	}

	if le.Client == nil {
		log.Println("No log for " + message)
		return nil
	}

	le.Client.Println(message)

	return nil
}

// LogInterface complex struct
func (le *LogEntry) LogInterface(i interface{}) error {

	le.Log(fmt.Sprintf("%#v", i))

	return nil
}

// LogErrInterface complex struct
func (le *LogEntry) LogErrInterface(i interface{}) error {

	le.Log(fmt.Sprintf("[Error] %#v", i))

	return nil
}
