package logger

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)
}

type DefaultLogger struct {
}

func (l *DefaultLogger) Debug(msg string) {
	debugWithSkip(2, msg)
}

func (l *DefaultLogger) Info(msg string) {
	infoWithSkip(2, msg)
}

func (l *DefaultLogger) Warn(msg string) {
	warnWithSkip(2, msg)
}

func (l *DefaultLogger) Error(msg string) {
	errorWithSkip(2, msg)
}

func (l *DefaultLogger) Fatal(msg string) {
	fatalWithSkip(2, msg)
}

func (l *DefaultLogger) Panic(msg string) {
	panicWithSkip(2, msg)
}
