package logging

type handler interface {
	record(level string, s interface{}, file string, line int)
}
