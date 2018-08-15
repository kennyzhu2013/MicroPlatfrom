// Package log is an interface for structured logging.
package log

import (
	"encoding/json"
)

const (
	DebugLevel Level = 0
	InfoLevel  Level = 1
	WarnLevel  Level = 2
	ErrorLevel Level = 3
	FatalLevel Level = 4
)

type Level int32

type Fields map[string]string

// A structure log interface which can
// output to multiple backends.
type Log interface {
	Init(opts ...Option) error
	Options() Options
	Logger
	String() string
}

type Logger interface {
	// Logger interface
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	// Formatted logger
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	// Specify your own levels
	Log(l Level, args ...interface{})
	Logf(l Level, format string, args ...interface{})
	// Returns with extra fields
	WithFields(f Fields) Logger
}

// Event represents a single log event
type Event struct {
	Level     Level  `json:"level"`
	Fields    Fields `json:"fields"`
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

// An output represents a file, indexer, syslog, etc
type Output interface {
	// Send an event
	Send(*Event) error
	// Flush any buffered events
	Flush() error
	// Discard the output
	Close() error
	// Name of output
	String() string
}

type Option func(o *Options)

type OutputOption func(o *OutputOptions)

var (
	DefaultLevel      Level = InfoLevel
	DefaultOutputName       = "log.json"

	Levels = map[Level]string{
		DebugLevel: "debug",
		InfoLevel:  "info",
		WarnLevel:  "warn",
		ErrorLevel: "error",
		FatalLevel: "fatal",
	}
)

func (e *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Level     string `json:"level"`
		Fields    Fields `json:"fields"`
		Message   string `json:"message"`
		Timestamp int64  `json:"timestamp"`
	}{
		Level:     Levels[e.Level],
		Fields:    e.Fields,
		Message:   e.Message,
		Timestamp: e.Timestamp,
	})
}

func NewLog(opts ...Option) Log {
	return newOS(opts...)
}
