package logging

type handler interface {
	record(level string, s string, file string, line int)
}
