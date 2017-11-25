package logger

import (
	"errors"
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

	// Connect the logger
	Connect(Settings)

	// Disconnect the logger
	Close() error

	// Set a Prefix Keyword to separate your log
	SetPrefix(string)

	// Share clones the current logger
	Share() Logger
}

// LogEntry contains requirements to log
type LogEntry struct {
	Logger
	Client *le_go.Logger
	Settings
	prefix string
}

// Settings contains credential
type Settings struct {
	State   bool   `yaml:"state"`
	Key     string `yaml:"key"`
	Verbose bool   `yaml:"verbose"`
}

// Check settings integrity
func (s Settings) Check() error {
	var err error

	if s.State == true && len(s.Key) == 0 {
		err = errors.New("Logger Key is required because State value is 'true'")
	}

	return err
}

// NewLogger constructor
func NewLogger() Logger {
	return &LogEntry{
		Settings: Settings{
			Verbose: true,
		},
	}
}

// Share allows to clone this object and use the same client and property for
// another kind of logger
func (le *LogEntry) Share() Logger {
	return &LogEntry{
		Settings: Settings{
			State:   le.State,
			Verbose: le.Verbose,
		},
		Client: le.Client,
		prefix: le.prefix,
	}
}

// Connect connect to LogEntries server
func (le *LogEntry) Connect(s Settings) {
	le.Settings = s
	if false == s.State {
		le.Log("Logger is disabled")
		return
	}

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

// Close disconnects logger
func (le *LogEntry) Close() error {
	if le.Client != nil {
		le.Client.Close()
	}

	return nil
}

// Log simple string
func (le *LogEntry) Log(message string) error {
	if le.Verbose {
		log.Printf("[%s] %s", le.prefix, message)
	}

	if false == le.State {
		return nil
	}

	if le.Client == nil {
		log.Println("[%s] No log for %s"+le.prefix, message)
		return nil
	}

	le.Client.Printf("[%s] %s", le.prefix, message)

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

// SetPrefix to the current loggger
func (le *LogEntry) SetPrefix(p string) {
	le.prefix = p
}
