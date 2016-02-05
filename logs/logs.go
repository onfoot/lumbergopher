package logs

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Warn
	Debug
	Error
	Other
)

func (l LogLevel) String() string {
	switch l {
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

func levelFromString(level string) (LogLevel, error) {
	switch level {
	case "Info":
		return Info, nil
	case "Warn":
		return Warn, nil
	case "Error":
		return Error, nil
	case "Debug":
		return Debug, nil
	}

	return Other, nil
}

// UnmarshalJSON parses a log level given as string in json
func (l *LogLevel) UnmarshalJSON(data []byte) error {

	dataString := string(data)
	dataString = strings.Trim(dataString, "\"")

	level, err := levelFromString(dataString)

	if err != nil {
		return &json.UnsupportedValueError{Value: reflect.ValueOf(data), Str: dataString}
	}

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
