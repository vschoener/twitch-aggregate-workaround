package logger

import (
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

	Connect(Settings)
}

// LogEntry contains requirements to log
type LogEntry struct {
	Logger
	Client *le_go.Logger
}

// Settings contains credential
type Settings struct {
	Key string `yaml:"key"`
}

// NewLogger constructor
func NewLogger() Logger {
	return &LogEntry{}
}

// Connect connect to LogEntries server
func (le *LogEntry) Connect(s Settings) {

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

	if le.Client == nil {
		log.Println("No log for " + message)
		return nil
	}

	le.Client.Println(message)

	return nil
}

// LogInterface complex struct
func (le *LogEntry) LogInterface(i interface{}) error {
	if le.Client == nil {
		return nil
	}

	le.Log(i.(string))

	return nil
}
