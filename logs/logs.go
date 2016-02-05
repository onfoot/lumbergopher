package logs

import (
	"strings"
	"time"
)

type LogLevel int

const (
	Verbose LogLevel = iota
	Info
	Warn
	Debug
	Error
	Other
)

func (l LogLevel) String() string {
	switch l {
	case Verbose:
		return "Verbose"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Debug:
		return "Debug"
	case Error:
		return "Error"
	}

	return "Other"
}

func levelFromString(level string) LogLevel {
	switch level {
	case "Verbose":
		return Verbose
	case "Info":
		return Info
	case "Warn":
		return Warn
	case "Error":
		return Error
	case "Debug":
		return Debug
	}

	return Other
}

// UnmarshalJSON parses a log level given as string in json
func (l *LogLevel) UnmarshalJSON(data []byte) error {

	dataString := string(data)
	dataString = strings.Trim(dataString, "\"")

	level := levelFromString(dataString)

	*l = level
	return nil
}

type LogEvent struct {
	Message string
	Level   LogLevel
	Time    time.Time
}

type LogEvents struct {
	Events []LogEvent
}
