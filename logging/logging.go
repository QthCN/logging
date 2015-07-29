package logging

import (
	"runtime"
)

type logger struct {
	name     string
	parent   *logger
	handlers []handler
}

var rootLogger logger
var defaultLoggerLevel string

func init() {
	rootLogger = logger{
		"rootLogger",
		nil,
		make([]handler, 0, 5),
	}

	defaultLoggerLevel = "debug"

	rootLogger.handlers = append(rootLogger.handlers, &SimpleHandler{})
}

func SetLoggerLevel(level string) {
	switch level {
	case "debug":
		fallthrough
	case "DEBUG":
		defaultLoggerLevel = "debug"
	case "info":
		fallthrough
	case "INFO":
		defaultLoggerLevel = "info"
	case "warn":
		fallthrough
	case "WARN":
		defaultLoggerLevel = "warn"
	case "error":
		fallthrough
	case "ERROR":
		defaultLoggerLevel = "error"
	default:
		defaultLoggerLevel = "debug"
	}
}

func GetLogger(name string) logger {
	l := logger{
		name,
		&rootLogger,
		make([]handler, 0, 5),
	}

	return l
}

func (l *logger) record(level string, s string, file string, line int) {
	for _, h := range l.handlers {
		h.record(level, s, file, line)
	}
	if l.parent != nil {
		l.parent.record(level, s, file, line)
	}
}

func (l *logger) Debug(s string) {
	_, file, line, _ := runtime.Caller(1)
	switch defaultLoggerLevel {
	case "debug":
		{
			l.record("debug", s, file, line)
		}
	}
}

func (l *logger) Info(s string) {
	_, file, line, _ := runtime.Caller(1)
	switch defaultLoggerLevel {
	case "debug":
		fallthrough
	case "info":
		{
			l.record("info", s, file, line)
		}
	}
}

func (l *logger) Warn(s string) {
	_, file, line, _ := runtime.Caller(1)
	switch defaultLoggerLevel {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "warn":
		{
			l.record("warn", s, file, line)
		}
	}
}

func (l *logger) Error(s string) {
	_, file, line, _ := runtime.Caller(1)
	switch defaultLoggerLevel {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "warn":
		fallthrough
	case "error":
		{
			l.record("error", s, file, line)
		}
	}
}
