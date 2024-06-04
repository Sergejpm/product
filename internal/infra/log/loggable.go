package log

type Loggable interface {
	Infof(template string, args ...any)
	Debugf(template string, args ...any)
	Warnf(template string, args ...any)
	Errorf(template string, args ...any)
	Fatalf(template string, args ...any)
	Info(text string)
	Debug(text string)
	Warn(text string)
	Error(text string)
	Fatal(text string)
}
