package main

import "time"

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

type LogEvent struct {
	Message string
	Level   LogLevel
	Time    time.Time
}

type LogEvents struct {
	Events []LogEvent
}
